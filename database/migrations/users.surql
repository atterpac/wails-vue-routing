DEFINE TABLE user;

DEFINE FIELD firstname ON TABLE user TYPE option<string>;
DEFINE FIELD lastname ON TABLE user TYPE option<string>;
DEFINE FIELD email ON TABLE user TYPE option<string>
    ASSERT string::is::email($value);
DEFINE FIELD username ON TABLE user TYPE string
    VALUE string::lowercase($value);
DEFINE FIELD password ON TABLE user TYPE string
    VALUE crypto::argon2::generate($value);
--DEFINE FIELD events ON TABLE user TYPE option<array>;
--DEFINE FIELD events.* ON TABLE user TYPE record();


--Modifications
DEFINE FIELD modified_at ON TABLE user TYPE option<datetime>
    VALUE time::now();
DEFINE FIELD created_at ON TABLE user TYPE option<datetime>;
DEFINE FIELD deleted_at ON TABLE user TYPE option<datetime>;
