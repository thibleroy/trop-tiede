#include <WiFi.h>
#include <WiFiMulti.h> 
#include <PubSubClient.h>
#include "DHT.h" 
//WIFI
const char* ssid = "Livebox-8521";
const char* password = "24021963";
//MQTT
#define MQTT_BROKER       "192.168.1.53"
#define MQTT_BROKER_PORT  1883
#define MQTT_USERNAME     "thibleroy"
#define MQTT_KEY          "thib" 
WiFiMulti WiFiMulti;
WiFiClient espClient;
PubSubClient client(espClient);
//DHT22 sensor
DHT DHT_sens(32, DHT22); // datapin sensor connected to pin 10 Arduino
void setup() {
  Serial.begin(115200);
  setup_wifi();
  setup_mqtt();
  client.publish("#", "Hello from ESP32");
  DHT_sens.begin();
 
}
void loop() {
  reconnect();
  client.loop(); 
  mqtt_publish("/room1/temperature", get_temp());
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

float h, t;
h = DHT_sens.readHumidity();
t = DHT_sens.readTemperature();

delay (2000); // pause a second
Serial.print ("Humidity: ");
Serial.print (h,0); // zero decimal
Serial.print (" %\t");
Serial.print ("Temperature: ");
Serial.print (t,1); // one decimal
Serial.println (" *C");
delay (2000); 
return t;
}
