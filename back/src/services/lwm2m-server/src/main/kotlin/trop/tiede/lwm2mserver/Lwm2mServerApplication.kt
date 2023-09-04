package trop.tiede.lwm2mserver

import com.fasterxml.jackson.core.JsonProcessingException
import org.eclipse.leshan.core.node.TimestampedLwM2mNodes
import org.eclipse.leshan.core.observation.CompositeObservation
import org.eclipse.leshan.core.observation.Observation
import org.eclipse.leshan.core.observation.SingleObservation
import org.eclipse.leshan.core.request.ReadRequest
import org.eclipse.leshan.core.request.SendRequest
import org.eclipse.leshan.core.response.ObserveCompositeResponse
import org.eclipse.leshan.core.response.ObserveResponse
import org.eclipse.leshan.core.response.ReadResponse
import org.eclipse.leshan.server.LeshanServer
import org.eclipse.leshan.server.LeshanServerBuilder
import org.eclipse.leshan.server.californium.endpoint.CaliforniumServerEndpointsProvider
import org.eclipse.leshan.server.californium.endpoint.coap.CoapServerProtocolProvider
import org.eclipse.leshan.server.observation.ObservationListener
import org.eclipse.leshan.server.queue.PresenceListener
import org.eclipse.leshan.server.registration.Registration
import org.eclipse.leshan.server.registration.RegistrationListener
import org.eclipse.leshan.server.registration.RegistrationUpdate
import org.eclipse.leshan.server.send.SendListener
import org.springframework.boot.CommandLineRunner
import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.runApplication
import org.springframework.context.ApplicationContext
import org.springframework.context.annotation.Bean
import kotlin.system.exitProcess

@SpringBootApplication
class Lwm2mServerApplication {
    lateinit var server: LeshanServer

    @Bean
    fun commandLineRunner(ctx: ApplicationContext): CommandLineRunner? {
        return CommandLineRunner {
            try {
                this.init()
                this.server.start()
            } catch (e: Exception) {
                print("Error on init/start LwM2M server. Error is: ${e.message}")
                exitProcess(1)
            }
        }
    }

    fun init() {
        val builder = LeshanServerBuilder()
        val endpointsBuilder = CaliforniumServerEndpointsProvider.Builder( // Add coap Protocol support
            CoapServerProtocolProvider()
        )  // Add COAPs protocol support
        builder.setEndpointsProvider(endpointsBuilder.build())
        this.server = builder.build()
        this.initServerObservationService()
        this.initServerRegistrationService()
        this.initServerSendService()
        this.initServerPresenceService()
        // Enregistrer un observateur pour les observations
    }

    fun initServerPresenceService() {
        this.server.presenceService.addListener(object : PresenceListener {
            override fun onAwake(registration: Registration?) {
                println("device is awake ${registration.toString()}")
            }

            override fun onSleeping(registration: Registration?) {
                println("device is sleeping ${registration.toString()}")
            }
        })
    }

    fun initServerSendService() {
        this.server.sendService.addListener(object : SendListener {
            override fun dataReceived(
                registration: Registration?,
                data: TimestampedLwM2mNodes?,
                request: SendRequest?
            ) {
                if (registration != null) {
                    try {
                        for (node in data!!.nodes) {
                            println("*** DATA RECEIVED *** /")
                            when (node.value.id) {
                                5700 -> print("Sensor Value ${node.value}")
                                else -> { // Note the block
                                    print("data node value id not handled")
                                }
                            }
                            println("node ${node.value}")
                        }
                    } catch (e: JsonProcessingException) {
                        println("Error while processing json [${data.toString()}] : [${e.message}]")
                    }
                }

            }

            override fun onError(registration: Registration?, error: java.lang.Exception?) {
                if (error != null) {
                    print("sendService error is ${error.message}")
                    return
                }
                print("sendService registration is ${registration.toString()}")
            }


        })
    }

    fun initServerObservationService() {
        this.server.observationService.addListener(object : ObservationListener {
            override fun newObservation(observation: Observation?, registration: Registration?) {
                println("new obs $observation.id")
            }

            override fun cancelled(observation: Observation?) {
                println("new obs $observation.id")
            }

            override fun onResponse(
                observation: SingleObservation?,
                registration: Registration?,
                response: ObserveResponse?
            ) {
                println("new obs $observation.id")
            }

            override fun onResponse(
                observation: CompositeObservation?,
                registration: Registration?,
                response: ObserveCompositeResponse?
            ) {
                println("new obs $observation.id")
            }

            override fun onError(observation: Observation?, registration: Registration?, error: Exception?) {
                println("new obs $observation.id")
            }

        })
    }

    fun make(registration: Registration, path: String) {
        println("path $path")
        val readRequest = ReadRequest(path)
        println("serv ${this.server}")
        val readResponse: ReadResponse = this.server.send(registration, readRequest)
        println("reg $registration")
        println("serv ${this.server}")
    }

    fun readAll() {
        try {

            this.server.registrationStore.allRegistrations.forEach { registration ->
                run {
                    println("registration ${registration.id}")
                    make(registration, "/3303/0/5700")
                }
            }
        } catch (e: Exception) {
            println("error on readAll ${e.message}")
        }

    }

    fun initServerRegistrationService() {
        this.server.registrationService.addListener(object : RegistrationListener {
            override fun registered(
                registration: Registration, previousReg: Registration?,
                previousObsersations: Collection<Observation?>?
            ) {
                println("new device: ${registration.id}")
                server.registrationStore.addRegistration(registration)
                readAll()
            }

            override fun updated(update: RegistrationUpdate?, updatedReg: Registration, previousReg: Registration?) {
                println("device is still here: ${updatedReg.id}")
                readAll()
            }

            override fun unregistered(
                registration: Registration, observations: Collection<Observation?>?, expired: Boolean,
                newReg: Registration?
            ) {
                println("device left: " + registration.endpoint)
            }
        })
    }
}

fun main(args: Array<String>) {
    runApplication<Lwm2mServerApplication>(*args)
}
