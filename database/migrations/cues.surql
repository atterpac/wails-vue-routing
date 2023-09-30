DEFINE TABLE cues;

DEFINE FIELD id ON TABLE cues TYPE string;
DEFINE FIELD num ON TABLE cues TYPE number;
DEFINE FIELD part ON TABLE cues TYPE number;
DEFINE FIELD time_up ON TABLE cues TYPE number;
DEFINE FIELD time_down ON TABLE cues TYPE number;
DEFINE FIELD delay ON TABLE cues TYPE number;
DEFINE FIELD follow ON TABLE cues TYPE number;
DEFINE FIELD block ON TABLE cues TYPE number;
DEFINE FIELD page ON TABLE cues TYPE string;
DEFINE FIELD system ON TABLE cues TYPE string;
DEFINE FIELD measure ON TABLE cues TYPE string;
DEFINE FIELD description ON TABLE cues TYPE string;
DEFINE FIELD notes ON TABLE cues TYPE string;
DEFINE FIELD placement ON TABLE cues TYPE string;
DEFINE FIELD scene ON TABLE cues TYPE record();
DEFINE FIELD worknotes ON TABLE cues TYPE array;
DEFINE FIELD worknotes.* ON TABLE cues TYPE record();
