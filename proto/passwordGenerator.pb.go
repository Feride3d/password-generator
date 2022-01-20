// Protocol Buffers Version

syntax = "proto3";

// the code obtained from the protobuff is added to the passwordservice package

package passwordGenerator;

message PasswordRequest {
    string sample = 1;
}

message PasswordResponse {
    string password = 1;
}

/* service with a Generate method that receives messages of type Request 
from the client and returns message of type Response */

service PasswordGeneratorService {
    rpc Generate(PasswordRequest) returns (PasswordResponse) {}
}