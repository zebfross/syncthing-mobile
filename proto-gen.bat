cd proto
protoc -I .. -I . --go_out=.. lib\protocol\bep.proto lib\protocol\deviceid_test.proto
protoc -I .. -I . --go_out=.. lib\protocol\bep.proto lib\protocol\deviceid_test.proto
protoc -I .. -I . --go_out=.. lib\config\authmode.proto lib\config\blockpullorder.proto lib\config\config.proto lib\config\deviceconfiguration.proto lib\config\folderconfiguration.proto lib\config\foldertype.proto lib\config\guiconfiguration.proto lib\config\ldapconfiguration.proto lib\config\ldaptransport.proto lib\config\observed.proto lib\config\optionsconfiguration.proto lib\config\pullorder.proto lib\config\size.proto lib\config\tuning.proto lib\config\versioningconfiguration.proto
protoc -I .. -I . --go_out=.. lib\fs\copyrangemethod.proto lib\fs\types.proto
protoc -I .. -I . --go_out=.. lib\db\structs.proto
protoc -I .. -I . --go_out=.. lib\discover\local.proto
cd ..
