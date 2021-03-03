#include "DHT.h"

float h, t;
DHT DHT_sens(4, DHT22); // datapin sensor connected to pin 10 Arduino

 

void setup() {

DHT_sens.begin();

Serial.begin (115200);
Serial.println ("===============================================");
Serial.println ("Bare DHT22 temp-humidity sensor – March, 2021");
Serial.println ("===============================================");
Serial.println (" ");
}

 

void loop(){

// ==== read from buffer and display =========================

h = DHT_sens.readHumidity();
t = DHT_sens.readTemperature();

delay (2000); // pause a second
Serial.print ("Humidity: ");
Serial.print (h,0); // zero decimal
Serial.print (" %\t");
Serial.print ("Temperature: ");
Serial.print (t,1); // one decimal
Serial.println (" *C");
delay (2000); // pause a second

}
