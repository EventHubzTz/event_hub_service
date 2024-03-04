#!/bin/bash
#VARIABLES
rootPath=/go/testbed/event_hub_service/shellScripts;
rootCustomCommandsFolder=/etc/customCommands;
publicCommandsFolder=/usr/bin;
rootServicesPath=/go/testbed/event_hub_service/shellScripts/services;
systemFolderForServices=/etc/systemd/system;
projectFolderPath=/go/testbed/event_hub_service;
#01. NAVIGATE TO THE PROJECT PATH
cd $projectFolderPath || exit;
#02. UPDATE GOLANG PROJECT
rm -rf main
go get
go mod download
go build main.go
#03. MOVE TO ETC FILE WITH SYSTEM CONFIGURATIONS
cd /etc || exit;
#04. MAKE customCommands FOLDER
rm -rf $rootCustomCommandsFolder
mkdir customCommands;
#05. MOVE TO THE ROOT PROJECT
cd $rootPath || exit;
#06. COPY ALL SHELL SCRIPTS TO customCommands FOLDER
for FILE in "$rootPath"/*.sh
do
        cp "$FILE" $rootCustomCommandsFolder
done
#07. MOVE TO customCommands FILE WITH SYSTEM CONFIGURATIONS
cd "$rootCustomCommandsFolder" || exit;
#08. CHANGE PERMISSIONS OF ALL .sh FILES TO BE EXECUTABLE FILES
fileCounter=0;
echo "-----------------------------------";
echo "           CHMOD - START           ";
echo "-----------------------------------";
for FILE in "$rootCustomCommandsFolder"/*.sh
do
       ((fileCounter=fileCounter+1))
        echo $fileCounter" - $FILE";
        chmod -R  u+x "$FILE"
done
echo "-----------------------------------";
echo "            CHMOD - END            ";
echo "-----------------------------------";
#09. TEST IF ALL COMMANDS ARE EXECUTABLE
fileCounter=0;
for FILE in *.sh
do
       ((fileCounter=fileCounter+1))
        echo $fileCounter" - $FILE";
        #REMOVE IF THERE ARE COMMANDS IN /USR/BIN/FOLDER
        rm -rf $publicCommandsFolder""/"$FILE";
done
#10. COPY ALL SHELL SCRIPTS TO /usr/bin FOLDER
fileCounter=0;
echo "-----------------------------------";
echo "     COPY TO /USR/BIN - START      ";
echo "-----------------------------------";
for FILE in "$rootCustomCommandsFolder"/*.sh
do
        ((fileCounter=fileCounter+1))
        cp "$FILE" $publicCommandsFolder
        echo $fileCounter" - $FILE" " - Copied";
done
echo "-----------------------------------";
echo "      COPY TO /USR/BIN - END       ";
echo "-----------------------------------";
#11. MOVE TO /services FOLDER
cd $rootServicesPath || exit;
#12 MOVE ALL SERVICES TO OPERATING SYSTEM SERVICE FOLDERS
fileCounter=0;
currentDir=$(pwd);
if [ $rootServicesPath = "$currentDir" ]
 then
    echo "------------------------------------";
    echo "COPY TO /ETC/SYSTEMD/SYSTEM - START";
    echo "------------------------------------";
    for FILE in *.service
    do
       echo "___________________________________________________________________";
      #I. SET . AS A DELIMITER
      IFS='.'

      #II. READ THE SPLIT WORDS INTO AN ARRAY BASED ON . DELIMITER
      read -r -a serviceNames <<< "$FILE"

      echo "${serviceNames[0]}"
      #III. CHECK IF THE SERVICE IS ALREADY LOADED TO THE SYSTEM
        if sudo systemctl list-unit-files --type service | grep -F "${serviceNames[0]}";
         then
            echo "service ${serviceNames[0]} - existed!";
            sudo service "${serviceNames[0]}" stop
           echo "service ${serviceNames[0]} - Stopped!";
         else
            echo "service ${serviceNames[0]} - Not existed!";
       fi
      #III. DISABLE IF THERE IS SYMLINK FOR STARTING THE SERVICE ON BOOT EXISTED
      if [ -f $systemFolderForServices""/"$FILE" ];
        then
          #VII. DISABLE SYMLINK FOR STARTING THE SERVICE ON BOOT
          echo $systemFolderForServices""/"$FILE - Existed"
          sudo systemctl disable "${serviceNames[0]}"
          rm -rf $systemFolderForServices""/"$FILE"
          echo $systemFolderForServices""/"$FILE - Removed"
        else
          echo "service ${serviceNames[0]} - Not existed!";
      fi
      #V. COPY FILE TO SYSTEM SERVICES
      cp "$FILE" $systemFolderForServices
      #VI. RELOAD SERVICES
      sudo systemctl daemon-reload
      #VII. CREATE SYMLINK FOR STARTING THE SERVICE ON BOOT
      sudo systemctl enable "${serviceNames[0]}"
      echo $systemFolderForServices""/"$FILE - Created"
      #VIII. START SERVICE
      sudo service "${serviceNames[0]}" start;
      echo "Service - ${serviceNames[0]} - started"
      echo "___________________________________________________________________";
    done
    echo "------------------------------------";
    echo " COPY TO /ETC/SYSTEMD/SYSTEM - END ";
    echo "------------------------------------";
 else
  echo "------------------------------------------------------------";
  echo " COMMAND NOT NAVIGATED TO THE SHELL SCRIPTS SERVICES FOLDER ";
  echo "------------------------------------------------------------";
fi