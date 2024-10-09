#! /bin/bash

for FOLDER in */; do

    FOLDER=${FOLDER%/}
    

    CREATION_DATE=$(stat "$FOLDER" | grep "Birth" | awk '{print $2}')
    

    if [ -z "$CREATION_DATE" ]; then
        echo "Creation date not available for $FOLDER. Skipping."
        continue
    fi


    NEW_FOLDER="${FOLDER}-${CREATION_DATE}"
    
    mv "$FOLDER" "$NEW_FOLDER"
    
    echo "Renamed $FOLDER to $NEW_FOLDER"
done
