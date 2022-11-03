cd proto
set GOPATH=--go_out=paths=source_relative:../syncthing/

protoc -I .. -I . %GOPATH% protocol\bep.proto protocol\deviceid_test.proto
protoc -I .. -I . %GOPATH% protocol\bep.proto protocol\deviceid_test.proto
protoc -I .. -I . %GOPATH% config\authmode.proto config\blockpullorder.proto config\config.proto config\deviceconfiguration.proto config\folderconfiguration.proto config\foldertype.proto config\guiconfiguration.proto config\ldapconfiguration.proto config\ldaptransport.proto config\observed.proto config\optionsconfiguration.proto config\pullorder.proto config\size.proto config\tuning.proto config\versioningconfiguration.proto
protoc -I .. -I . %GOPATH% fs\copyrangemethod.proto fs\types.proto
protoc -I .. -I . %GOPATH% db\structs.proto
protoc -I .. -I . %GOPATH% discover\local.proto
cd ..
