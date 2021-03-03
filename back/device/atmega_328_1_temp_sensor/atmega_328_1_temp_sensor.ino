#include <dht11.h>                                              
#define DHT11PIN 2                                           
dht11 DHT11;                                                       
void setup() {                                                       
Serial.begin(9600);                                          
pinMode(DHT11PIN,OUTPUT);                      
}
void loop() {                                                 
int chk = DHT11.read(DHT11PIN);                 
int tem=(float)DHT11.temperature;              
int hum=(float)DHT11.humidity;                  
Serial.print("Tempeature:");                        
Serial.println(tem);                                   
Serial.print("Humidity:");                           
Serial.print(hum);                                    
Serial.println("%");                                  
delay(1000);                                          
}
