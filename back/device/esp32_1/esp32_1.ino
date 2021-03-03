#include <WiFi.h>
#include <WiFiMulti.h>


const char* ssid = "Livebox-8521";
const char* password = "24021963";

WiFiMulti WiFiMulti;


WiFiClient espClient;

void setup() {
  Serial.begin(115200);
  setup_wifi();
}

void loop() {
  
  }

void setup_wifi(){
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
