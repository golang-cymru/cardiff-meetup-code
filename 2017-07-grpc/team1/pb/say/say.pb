syntax = "proto3";

package say


service TextToSpeech {
    rpc SaySomething(Something) returns (Result);
}


message Something {
    string message = 1;
}


message Result {
    bytes audio
}

