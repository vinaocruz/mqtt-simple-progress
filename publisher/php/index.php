<?php

require_once __DIR__ . '/vendor/autoload.php';

use PhpMqtt\Client\MqttClient;

$mqtt = new MqttClient("localhost");
$mqtt->connect();
$mqtt->publish("progress", 15);
$mqtt->disconnect();
