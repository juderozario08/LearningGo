check='main.go'

THEN=$(stat -c %z ${check})
touch ${check}
NOW=$(stat -c %z ${check})

while true; do
    sleep 0.1
    if [ "$NOW" != "$THEN" ]; then
        clear
        go run .
        THEN=$NOW
        echo
    fi
    NOW=$(stat -c %z ${check})
done
