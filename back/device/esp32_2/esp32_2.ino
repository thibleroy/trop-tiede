#include <WiFi.h>
#include <WiFiMulti.h> 
#include <PubSubClient.h>
#include "DHT.h" 
//WIFI
const char* ssid = "XXXXXX";
const char* password = "*********";
//MQTT
#define MQTT_BROKER       "192.168.0.X"
#define MQTT_BROKER_PORT  1883
#define MQTT_USERNAME     "XXXXXXX"
#define MQTT_KEY          "*******" 
WiFiMulti WiFiMulti;
WiFiClient espClient;
PubSubClient client(espClient);
//DHT11 sensor
DHT DHT_sens(4, DHT11); // datapin sensor connected to pin 4 Arduino
void setup() {
  Serial.begin(115200);
  setup_wifi();
  setup_mqtt();
  client.publish("/room1", "Hello from ESP32");
  client.subscribe("/room1");
  DHT_sens.begin();
}
void loop() {
  reconnect();
  client.loop(); 
  mqtt_publish("/room1/temperature", get_temp());
  mqtt_publish("/room1/humidity", get_humidity());
  delay (10000);
}
void setup_wifi(){
  //connexion au wifi
  WiFiMulti.addAP(ssid, password);
  while ( WiFiMulti.run() != WL_CONNECTED ) {
    delay ( 500 );
    Serial.print ( "." );
  }
  Serial.println("");
  Serial.println("WiFi connecté");
  Serial.print("MAC : ");
  Serial.println(WiFi.macAddress());
  Serial.print("Adresse IP : ");
  Serial.println(WiFi.localIP());
}
void setup_mqtt(){
  client.setServer(MQTT_BROKER, MQTT_BROKER_PORT);
  reconnect();
}
void reconnect(){
  while (!client.connected()) {
    Serial.println("Connection au serveur MQTT ...");
    if (client.connect("ESPClient", MQTT_USERNAME, MQTT_KEY)) {
      Serial.println("MQTT connecté");
    }
    else {
      Serial.print("echec, code erreur= ");
      Serial.println(client.state());
      Serial.println("nouvel essai dans 2s");
    delay(2000);
    }
  }
}
//Fonction pour publier un float sur un topic
void mqtt_publish(String topic, float t){
  char top[topic.length()+1];
  topic.toCharArray(top,topic.length()+1);
  char t_char[50];
  String t_str = String(t);
  t_str.toCharArray(t_char, t_str.length() + 1);
  client.publish(top,t_char);
}

float get_temp(){

  float t;
  t = DHT_sens.readTemperature();

  Serial.print ("Temperature: ");
  Serial.print (t,1); // one decimal
  Serial.println (" *C");
  return t;
}

float get_humidity(){

  float h;
  h = DHT_sens.readHumidity();

  Serial.print ("Humidity: ");
  Serial.print (h,0); // zero decimal
  Serial.print (" %\t");
  return h;
}
