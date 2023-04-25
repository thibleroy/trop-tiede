plugins {
    kotlin("jvm") version "1.8.0"
    application
}

group = "trop.tiede"
version = "1.0.0"

repositories {
    mavenCentral()
}

dependencies {
    implementation("org.eclipse.leshan:leshan-server-cf:2.0.0-M10")
    implementation("org.slf4j:slf4j-simple:2.0.5")
    testImplementation("org.jetbrains.kotlin:kotlin-test:1.8.20-RC")
}

tasks.test {
    useJUnitPlatform()
}

kotlin {
    jvmToolchain(11)
}

application {
    mainClass.set("MainKt")
}