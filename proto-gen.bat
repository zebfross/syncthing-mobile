cd proto
protoc -I .. -I . --go_out=paths=source_relative:../syncthing/ protocol\bep.proto protocol\deviceid_test.proto
protoc -I .. -I . --go_out=paths=source_relative:../syncthing/ protocol\bep.proto protocol\deviceid_test.proto
protoc -I .. -I . --go_out=paths=source_relative:../syncthing/ config\authmode.proto config\blockpullorder.proto config\config.proto config\deviceconfiguration.proto config\folderconfiguration.proto config\foldertype.proto config\guiconfiguration.proto config\ldapconfiguration.proto config\ldaptransport.proto config\observed.proto config\optionsconfiguration.proto config\pullorder.proto config\size.proto config\tuning.proto config\versioningconfiguration.proto
protoc -I .. -I . --go_out=paths=source_relative:../syncthing/ fs\copyrangemethod.proto fs\types.proto
protoc -I .. -I . --go_out=paths=source_relative:../syncthing/ db\structs.proto
protoc -I .. -I . --go_out=paths=source_relative:../syncthing/ discover\local.proto
cd ..
