DEFINE TABLE user;

DEFINE FIELD id ON TABLE user TYPE string
    ASSERT string::is::uuid($value);
DEFINE FIELD firstname ON TABLE user TYPE string;
DEFINE FIELD lastname ON TABLE user TYPE string;
DEFINE FIELD email ON TABLE user TYPE string
    ASSERT string::is::email($value);
DEFINE FIELD username ON TABLE user TYPE string;
DEFINE FIELD password ON TABLE user TYPE string
    VALUE crypto::argon2::generate($value);
DEFINE FIELD events ON TABLE user TYPE array;
DEFINE FIELD events.* ON TABLE user TYPE record();


--Modifications
DEFINE FIELD modified_at ON TABLE user TYPE datetime;
DEFINE FIELD created_at ON TABLE user TYPE datetime;
DEFINE FIELD deleted_at ON TABLE user TYPE datetime;
