FROM jetbrains/qodana-jvm-android:2024.1


WORKDIR /app

COPY mqtt_cli_arm64 mqtt_cli_arm64
COPY conf.json conf.json

ENTRYPOINT [ "mqtt_cli_arm64" ]
