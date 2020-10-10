protoMongo - MongoDB
====================

The protoBSON package contains codecs for marshaling and unmarshalling common protobuf
data structures to more user-friendly BSON documents and values.

Basic Usage
-----------

Lets start with a protobuf file. At *Hogwarts School of Witchcraft and Wizardry*, after
the ``SortingHat`` service finishes sorting a new wizard, it seends a record to the
archives to be added to the school's roster. The archives stores the student roster in
MongoDB, so they need to serialize the incoming protobuf message to BSON.

.. code-block:: protobuf

  syntax = "proto3";
  package proto_cereal_doc;
  option go_package = "github.com/illuscio-dev/protoCereal-go/docs";

  import "cereal_proto/decimal.proto";
  import "cereal_proto/uuid.proto";
  import "cereal_proto/raw_data.proto";
  import "google/protobuf/wrappers.proto";
  import "google/protobuf/timestamp.proto";
  import "google/protobuf/any.proto";

  // The hogwarts houses.
  enum Houses {
    GRYFFINDOR = 0;
    RAVENCLAW = 1;
    HUFFLEPUFF = 2;
    SLYTHERIN = 3;
  }

  // Information about a Wizard's wand.
  message Wand {
    string core = 1;
  }

  // Information about a Wizard's sword.
  message Sword {
    string metal = 2;
  }

  // Information about a Hogwarts wizard.
  message Wizard {
    // The name of the Wizard.
    // @inject_tag: bson:"name"
    string name = 1;

    // A unique Identifier for the Wizard.
    // @inject_tag: bson:"id"
    cereal.UUID id = 2;

    // The exact moment the Wizard was sorted.
    // @inject_tag: bson:"sorted_at"
    google.protobuf.Timestamp sorted_at = 3;

    // The house this wizard belongs to.
    // @inject_tag: bson:"hogwarts_house"
    Houses hogwarts_house = 4;

    // The current balance of the wizard's Gringott's account in Galleons.
    // @inject_tag: bson:"gingotts_balance"
    cereal.Decimal gingotts_balance = 5;

    // Name of the Wizard's familiar. Nil if Wizard does not have a familiar.
    // @inject_tag: bson:"familiar_name"
    google.protobuf.StringValue familiar_name = 6;

    // Image of the Wizard.
    // @inject_tag: bson:"portrait"
    cereal.RawData portrait = 7;

    // The preferred weapon of this wizard.
    // @inject_tag: bson:"weapon"
    oneof weapon {
      Wand wand = 8;
      Sword sword = 9;
    }

    // An object the wizard is destined to obtain.
    // @inject_tag: bson:"destined_object"
    google.protobuf.Any destined_object = 10;
  }

.. note::

  We are injecting custom struct tags through
  `protoc-go-inject-tag <https://github.com/favadi/protoc-go-inject-tag>`_ for more
  human-friendly bson keys.

.. note::

  Hogwarts is a great place to work for code magicians! I highly recommend
  `this overview of magical tech interviews
  <https://aphyr.com/posts/341-hexing-the-technical-interview>`_
  to brush up on your skills prior to applying.

Back the problem at hand. Instantiating a new record in our Go code might look something
like this:

.. code-block:: go

  // Parse account balance as galleon decimal.
  gringottsBalance, err := primitive.ParseDecimal128("50625.56713")
  if err != nil {
    panic(fmt.Errorf("error parsing gringott's balance: %w", err))
  }

  // This student is destined to wield a sword, so pack that into an Any payload.
  destinedSword := &docs.Sword{Metal: "Silver"}
  destinedObject, err := anypb.New(destinedSword)
  if err != nil {
    panic(fmt.Errorf("error packing sword: %w", err))
  }

  // Create our wizard record.
  wizard := &docs.Wizard{
    Name:            "Harry Potter",
    Id:              messagesCereal.MustUUIDRandom(),
    SortedAt:        timestamppb.New(time.Now().UTC()),
    HogwartsHouse:   docs.Houses_GRYFFINDOR,
    GingottsBalance: messagesCereal.DecimalFromBson(gringottsBalance),
    FamiliarName:    &wrapperspb.StringValue{Value: "Hedwig"},
    Portrait:        &messagesCereal.RawData{Data: []byte("some image bytes")},
    Weapon: &docs.Wizard_Wand{
      Wand: &docs.Wand{
        Core: "Phoenix Feather",
      },
    },
    DestinedObject: destinedObject,
  }

To naively marshall this to BSON, we make the following incantation:

.. code-block:: go

  // Marshall to BSON binary.
  encoded, err := bson.Marshal(wizard)
  if err != nil {
    panic(fmt.Errorf("error marshalling wizard: %w", err))
  }

  // Extract into a bson map so we can peer into the document's structure.
  rawDocument := bson.M{}
  err = bson.Unmarshal(encoded, rawDocument)
  if err != nil {
    panic(fmt.Errorf("error unmarshalling to bson.M: %w", err))
  }

  // Pretty print the document
  _, _ = pretty.Println(rawDocument)

.. note::

  I am using `pretty <github.com/kr/pretty>`_ here to do the printing.

Which gives us the following output:

.. code-block:: text

  primitive.M{
      "id": primitive.M{
          "value": primitive.Binary{
              Subtype: 0x0,
              Data:    {0x46, 0xf3, 0x74, 0xc5, 0x14, 0x9f, 0x4c, 0x6a, 0x97, 0xad, 0xa7, 0xcd, 0xf3, 0x54, 0xb8, 0xa0},
          },
      },
      "sorted_at": primitive.M{
          "seconds": int64(1600743664),
          "nanos":   int32(229350000),
      },
      "hogwarts_house":   int32(0),
      "gingotts_balance": primitive.M{
          "high": int64(3473964162562916352),
          "low":  int64(5062556713),
      },
      "weapon": primitive.M{
          "wand": primitive.M{
              "core": "Phoenix Feather",
          },
      },
      "destined_object": primitive.M{
          "typeurl": "type.googleapis.com/proto_cereal_doc.Sword",
          "value":   primitive.Binary{
              Subtype: 0x0,
              Data:    {0x12, 0x6, 0x53, 0x69, 0x6c, 0x76, 0x65, 0x72},
          },
      },
      "name":     "Harry Potter",
      "portrait": primitive.M{
          "data": primitive.Binary{
              Subtype: 0x0,
              Data:    {0x73, 0x6f, 0x6d, 0x65, 0x20, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x20, 0x62, 0x79, 0x74, 0x65, 0x73},
          },
      },
      "familiar_name": primitive.M{
          "value": "Hedwig",
      },
  }

There are a lot of issues with this document, most of which involve unnecessary nesting:

  - Since optional values are put into wrappers, ``"familiar_name"`` is an object with a
    ``"value"`` sub-field, rather than a direct, nullable string value (what the
    protobuf data model actually represents).

  - This also happens with ``"portrait"`` and ``"id"`` fields, as both are just wrapper
    messages around a ``bytes`` field used to denote the kind of value these bytes
    represent. We don't need this extra level of nesting in our db.

  - BSON has native ``bsontype.BinaryUUID`` subtype for binary data, which is not being
    used on the ``"id"`` field.

  - The ``"gingotts_balance"`` field is encoded to a sub-document rather than a native
    ``primitive.Decimal128`` value.

  - Likewise, the ``"sorted_at"`` field is not being encoded into a
    ``primitive.DateTime`` value, even though that is the datatype
    ``timestamppb.Timestamp`` represents.

  - Because of the way oneof fields are implemented in Go, the ``"weapon"`` field has
    a nested ``"wand"`` document. If we instead had a ``doc.Sword`` message, we would
    have a nested ``"sword"`` field.

  - ``"hogwarts_house"`` is an inscrutable int value, making the document less friendly
    to human eyes when paging through the database directly.

  - ``"destined_object"`` is serialized as a binary blob, rather than a sub-document
    with fields we can query and inspect.

The nesting makes reasoning about and querying our data structure more complicated than
it needs to be.

Additionally, encoding this way has some unintuitive gotchas. The familiar's name is
found at ``"familiar_name.value"`` if it has one, but ``"familiar_name"`` is the field
that is nulled if no familiar exists.

We COULD define some custom structs to transfer the data to before marshalling, but in
addition to having to do a full copy of our data, this also adds an additional layer
of data modeling we need to keep in sync for our project. That MAY be good for some
larger projects, but it's a headache for prototyping and rapid development.

BSON `uses a registry <https://godoc.org/go.mongodb.org/mongo-driver/bson/bsoncodec>`_
to lookup type-based encoders and decoders when marshalling and unmarshalling structs.
Let's create a custom registry with
`ValueCodecs <https://godoc.org/go.mongodb.org/mongo-driver/bson/bsoncodec#ValueCodec>`_
registered to handle these proto message types:

.. code-block:: go

  // Create a new BSON registry builder
  registryBuilder := bson.NewRegistryBuilder()

  // Hand it off to protoBson to register our new codecs. We don't need to pass an
  // options object if we just want the default options.
  err = protoBson.RegisterCerealCodecs(registryBuilder, nil)
  if err != nil {
    panic(fmt.Errorf("error building registry: %w", err))
  }

  // Build our registry
  registry := registryBuilder.Build()

  // marshall to bson bytes with our new registry.
  encoded, err := bson.MarshalWithRegistry(registry, wizard)
  if err != nil {
    panic(fmt.Errorf("error marshalling with registry: %w", err))
  }

  // Extract into a bson map to see how this changed the document.
  rawDocument := bson.M{}
  err = bson.Unmarshal(encoded, rawDocument)
  if err != nil {
    panic(fmt.Errorf("error unmarshalling to bson.M: %w", err))
  }

  _, _ = pretty.Println(rawDocument)

Output:

.. code-block:: text

  primitive.M{
      "name": "Harry Potter",
      "id":   primitive.Binary{
          Subtype: 0x4,
          Data:    {0xf4, 0x10, 0x3, 0x6b, 0xa9, 0x67, 0x47, 0xeb, 0x95, 0xd2, 0x68, 0xa5, 0x61, 0x94, 0x53, 0x8a},
      },
      "sorted_at":        primitive.DateTime(1600748732738),
      "hogwarts_house":   int32(0),
      "gingotts_balance": primitive.Decimal128{h:0x3036000000000000, l:0x12dc07c29},
      "familiar_name":    "Hedwig",
      "portrait":         primitive.Binary{
          Subtype: 0x80,
          Data:    {0x73, 0x6f, 0x6d, 0x65, 0x20, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x20, 0x62, 0x79, 0x74, 0x65, 0x73},
      },
      "weapon": primitive.M{
          "wand": primitive.M{
              "core": "Phoenix Feather",
          },
      },
      "destined_object": primitive.M{
          "pb_type": "type.googleapis.com/proto_cereal_doc.Sword",
          "metal":   "Silver",
      },
  }

This looks much better! ``"familiar_name"`` has been extracted from it sub-document into
a plain string field (which it represents), the ``"id"`` and ``"portrait"`` fields are
no longer nested, and are of the correct binary subtype, the ``"sorted_at"`` field
now contains a BSON ``primitive.DateTime``, and ``"gingotts_balance"`` contains a native
``primitive.Decimal128``. ``"destined_object"`` is now a scrutable sub-document rather
than an opaque binary blob.

If we unmarshall the document, our data should re-populate into it's protobuf message
structure:

.. code-block:: go

  decoded := new(docs.Wizard)
  err = bson.UnmarshalWithRegistry(registry, encoded, decoded)
  if err != nil {
    panic(fmt.Errorf("error unmarshalling to struct: %w", err))
  }

BUT, we get an error! 🐞

.. code-block:: text

  panic: error unmarshalling to struct: error decoding key weapon: no decoder found for docs.isWizard_Weapon

Uh oh! The ``oneof Weapon weapon`` field is not being decoded properly. There is a way
to fix that! We'll come back to it in a moment.

Default Codecs
--------------

Registering the default options adds codecs that convert between the following
protobuf and BSON types:

==========================    ==========================  ==========================
Proto Type                    BSON Type                   Binary Subtype
==========================    ==========================  ==========================
\*wrapperspb.BoolValue        boolean
\*wrapperspb.BytesValue       primitive.Binary            bsontype.BinaryGeneric
\*wrapperspb.FloatValue       double
\*wrapperspb.DoubleValue      double
\*wrapperspb.Int32Value       int32
\*wrapperspb.Int64Value       int32 / int64
\*wrapperspb.StringValue      string
\*wrapperspb.UInt32Value      int32 / int64
\*wrapperspb.UInt64Value      int32 / int64
\*anypb.Any                   bsontype.EmbeddedDocument
\*timestamppb.Timestamp       primitive.DateTime
\*messagesCereal.Decimal      primitive.Decimal128
\*messagesCereal.RawData      primitive.Binary,           bsontype.BinaryUserDefined
\*messagesCereal.UUID         primitive.Binary,           bsontype.UUID
==========================    ==========================  ==========================

.. note::

  The above packages ``wrappers``, ``timestamppb``, and ``anypb`` are all from the
  `Go implementation <https://godoc.org/google.golang.org/protobuf/types/known/wrapperspb>`_ of google's
  `well known types <https://developers.google.com/protocol-buffers/docs/reference/google.protobuf>`_
  protobuf definitions.

.. note::

  All above codecs also support null values which encode from nil pointers to
  ``bsontype.Null`` and vice-versa.

.. warning::

  BSON only has int32 and int64 values. uint values which overflow an int64 will result
  in a marshalling error.

  Likewise, bson only allows for double-precision floats, so errors may occur from
  constant round-tripping of single precision floating-point values.

Any Fields
----------

``*anypb.Any`` fields are encoded by adding the type url as a field to an embedded
document created from the payload.

Because of implementation details, the internal payload must be serialized,
de-serialized, then serialized again to add the type url. There may be performance
implications for large messages stored in an ``*anypb.Any`` field.

Oneof Fields
------------

We still have to resolve the error caused by our oneof field. Under the hood, oneof
values are represented by a shared interface that a wrapper struct for each possible
value implements.

In our case, the interface is ``docs.isWizard_Weapon``. BSON has no idea what to do when
asked to decode into a literal interface, so we need to register a codec for it.

Auto-Generation
###############

Luckily, with protoCereal creating a codec for oneof fields is easy as:

.. code-block:: go

  // Create a new registry builder
  registryBuilder := bson.NewRegistryBuilder()

  // Create new options object
  cerealOpts := protoBson.NewMongoOpts().
    // passing one or more message types to this option automatically finds and
    // creates codecs for all oneof fields they contain.
    WithOneOfFields(new(docs.Wizard))

  // Hand it off to protoBson to register our new codecs.
  err = protoBson.RegisterCerealCodecs(registryBuilder, cerealOpts)
  if err != nil {
    panic(fmt.Errorf("error building registry: %w", err))
  }

  // Build our registry
  registry := registryBuilder.Build()

Now when we marshal the document, we get:

.. code-block:: text

  primitive.M{
      ...
      "weapon": primitive.M{
          "pb_type": "proto_cereal_doc.Wand",
          "core":    "Phoenix Feather",
      },
      ...
  }

Unmarshalling the document to a struct now works correctly as well.

.. note::

  Oneof codecs must be able to generate a 1-1 relationship between the encoded BSON
  type and it's decoded protobuf counterpart. Codec creation will fail if a oneof field
  contains two types that both map to the same BSON type. For instance if a oneof
  contains more than one of either ``float``, ``double``, ``wrappers.FloatValue``,
  or ``wrappers.DoubleValue``, the codec builder will panic while calling the
  ``protoBson.MongoOpts.WithOneOfFields`` setting.

  We panic because it is ambiguous as to what sort of value we should decode the raw
  ``bsontype.Double`` to when encountering it in a document.

  The exception is messages values which encode to ``bsontype.EmbeddedDocument``, for
  which we can inject the protobuf type name to aid in correct type-resolution on the
  way out, much like we do for ``anypb.Any``.

  Oneof interfaces that have multiple entries which encode to the same BSON type are
  currently outside the scope of this library, and will need to have hand-written codecs
  created for them.

.. warning::

  In the above example, the message type is encoded in the added ``"pb_type"`` field.
  This occurs when a oneof field has two possible types that both encode to embedded
  documents. If there is only one type the encodes to an embedded document, the type
  information is omitted.

  Because of certain implementation constraints, adding this type information requires
  encoding the object, decoding it, and encoding it again with the added field, which
  may have performance implications for large data structures.

Custom BSON Mapping
###################

It may be the case that protoCereal cannot deduce correctly what a oneof wrapper type
will encode to / decode from when round-tripping through BSON. We can supply our
own inner value type mapping in these cases.

Lets say we have the following protobuf file:

.. code-block:: protobuf

  import "cereal_proto/decimal.proto";

  message DecimalList {
    // @inject_tag: bson:"value"
    repeated cereal.Decimal value = 1;
  }

  message HasCustomOneOf {
    // @inject_tag: bson:"many"
    oneof many {
      cereal.Decimal decimal_value = 1;
      string string_value = 2;
      DecimalList decimal_list = 3;
    }
  }

We have registered ``DecimalList`` as a custom wrapper type (see
`Custom Wrapper Types`_):

.. code-block:: go

    cerealOpts := protoBson.
      NewMongoOpts().
      WithCustomWrappers(
        new(messagesCereal_test.DecimalList),
      ).
      WithOneOfFields(new(messagesCereal_test.HasCustomOneOf))

When protoCereal attempts to encode the ``decimal_list`` subfield of the ``many``
oneof, it will not realize that ``messagesCereal_test.DecimalList`` will encode
to an array of decimal values, rather than an embedded document.

We can signal that this will be the case like so:

.. code-block:: go

  cerealOpts := protoBson.
    NewMongoOpts().
    WithCustomWrappers(
      new(messagesCereal_test.DecimalList),
    ).
    WithOneOfElementMapping(
      new(messagesCereal_test.DecimalList),
      bsontype.Array,
      0x0,
    ).
    WithOneOfFields(new(messagesCereal_test.HasCustomOneOf))

Now protoCereal knows that to interpret arrays as ``*messagesCereal_test.DecimalList``
so whenever a oneof wrapper is wrapping this type, the correct mapping will be used.

Enums
-----

We may want to store enums as their string representation in our database for more
human-friendly documents. By default, proto enums are encoded to BSON as
``bsontype.Int32``. However, we can change that with the following option:

.. code-block:: go

  // Create a new options object
  cerealOpts := protoBson.NewMongoOpts().
    // Enable conversion of enums into strings for all enum types
    WithEnumStrings(true)

Now when we encode our document, we get:

.. code-block:: text

  primitive.M{
    ...
    "hogwarts_house":   "GRYFFINDOR",
    ...
  }

.. important::

  This option turns on string conversion for ALL proto enum types. protoCereal
  currently does not support the selective conversion of enum types, but such a feature
  could be added if there is interest!

Custom Wrapper Types
--------------------

Google uses wrapper types from it's
`wrappers from it's Well Known Types package <https://godoc.org/google.golang.org/protobuf/types/known/wrapperspb>`_
to represent nillable values. Using the default options, protoCereal will
register all of the wrapper types from google's Well Known Types into the codec
registry.

You can register custom wrapper types as well. For instance, google does not have
a defined wrapper for a fixed64 value, we can define one like so:

.. code-block:: protobuf

  // Wrapper for fixed int64.
  message Fixed64Value {
    fixed64 value= 1;
  }

Register the message as a wrapper type like so:

.. code-block:: go

  // Create a registry builder
  registryBuilder := bsoncodec.NewRegistryBuilder()

  // Pass an empty proto message to the WithCustomWrappers option to add it's type
  // as a wrapper type.
  opts := protoBson.NewMongoOpts().
    WithCustomWrappers(new(docsProto.Fixed64Value))

  // Register protoCereal BSON codecs with the registry.
  err := protoBson.RegisterCerealCodecs(registryBuilder, opts)
  if err != nil {
    panic(fmt.Errorf("error adding to registry: %w", err))
  }

  // Build the registry
  registry := registryBuilder.Build()

Now we can use it while serializing a struct to extract the inner value directly
into the key holding our wrapper:

.. code-block:: go

  type hasWrapper struct {
    Info *docsProto.Fixed64Value
  }

  message := &hasWrapper{
    Info: &docsProto.Fixed64Value{Value: 123456789},
  }

  serialized, err := bson.MarshalWithRegistry(registry, message)
  if err != nil {
    panic(fmt.Errorf("error serializing message: %w", err))
  }

  document := bson.M{}
  err = bson.UnmarshalWithRegistry(registry, serialized, document)
  if err != nil {
    panic(fmt.Errorf("error umarshalling to BSON document: %w", err))
  }

  fmt.Println("BSON DOCUMENT:", document)

  unmarshalled := new(hasWrapper)
  err = bson.UnmarshalWithRegistry(registry, serialized, unmarshalled)
  if err != nil {
    panic(fmt.Errorf("error unmarhalling to struct: %w", err))
  }

  fmt.Printf("\nUNMARSHALLED: %+v\n", unmarshalled)

Output:

.. code-block:: text

  BSON DOCUMENT: map[info:123456789]

  UNMARSHALLED: &{Info:value:123456789}

A wrapper type must:

- Be a pointer to a struct.

- Implement the ``proto.Message`` interface.

- Contain only one public field.

- The field must be called "Value"

The codec will serialize a ``bsontype.Null`` value if the wrapper message is a nil
pointer.
