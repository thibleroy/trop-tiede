#include <LiquidCrystal.h>  // blibliotheque LiquidCrystal

// définition des broches auxquelles on a branché l'afficheur LCD
const int pinRS = 15;      // broche 4 (RS) de l'afficheur branchée à GPIO04 de l'ESP8266
const int pinEnable = 13;  // broche 6 (Enable) de l'afficheur branchée à GPIO05 de l'ESP8266
const int pinD4 = 12;  // broche 11 (D4) de l'afficheur branchée à GPIO12 de l'ESP8266
const int pinD5 = 14;  // broche 12 (D5) de l'afficheur branchée à GPIO13 de l'ESP8266
const int pinD6 = 26;  // broche 13 (D6) de l'afficheur branchée à GPIO14 de l'ESP8266
const int pinD7 = 25;  // broche 14 (D7) de l'afficheur branchée à GPIO15 de l'ESP8266

// Initialisation de la bibliothèque LiquidCrystal en utilisant les broches définies ci-dessus:
LiquidCrystal lcd(pinRS, pinEnable, pinD4, pinD5, pinD6, pinD7);

void setup() {
  // On indique que notre afficheur comporte 2 lignes de 16 caractères:
  lcd.begin(16, 2);
}


void loop() {
  lcd.clear(); // on efface tout
  lcd.print("  Electronique");
  lcd.setCursor(0, 1); // deuxième ligne
  lcd.print("   en amateur");
  delay(5000);

  lcd.clear(); // on efface tout
  lcd.print("LCD et");
  lcd.setCursor(0, 1); // deuxième ligne
  lcd.print("       ESP8266");
  delay(5000);
}
