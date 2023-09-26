DEFINE TABLE events SCHEMAFULL;

--Define Feilds
DEFINE FIELD id ON TABLE events TYPE string
    ASSERT string::is::uuid($value);
DEFINE FIELD name ON TABLE events TYPE string;

--Users able to access the event
DEFINE FIELD users ON TABLE events TYPE array;
DEFINE FIELD users.* ON TABLE events TYPE record();

--Cue lists
DEFINE FIELD cue_lists ON TABLE events TYPE array;
DEFINE FIELD cue_lists.* ON TABLE events TYPE record();

--Followspots
DEFINE FIELD followspots ON TABLE events TYPE array;
DEFINE FIELD followspots.* ON TABLE events TYPE record();

--Worknotes
DEFINE FIELD worknotes ON TABLE events TYPE array;
DEFINE FIELD worknotes.* ON TABLE events TYPE record();

--Modifications
DEFINE FIELD modified_at ON TABLE user TYPE datetime;
DEFINE FIELD created_at ON TABLE user TYPE datetime;
DEFINE FIELD deleted_at ON TABLE user TYPE datetime;
