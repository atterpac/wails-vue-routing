DEFINE TABLE cue_lists;

DEFINE FIELD id ON TABLE cue_lists TYPE string
    ASSERT string::is::uuid(id);
DEFINE FIELD name ON TABLE cue_lists TYPE string;
DEFINE FIELD numb ON TABLE cue_lists TYPE number;
DEFINE FIELD description ON TABLE cue_lists TYPE string;
DEFINE FIELD cues ON TABLE cue_lists TYPE array;
DEFINE FIELD cues.* ON TABLE cue_lists TYPE record();

