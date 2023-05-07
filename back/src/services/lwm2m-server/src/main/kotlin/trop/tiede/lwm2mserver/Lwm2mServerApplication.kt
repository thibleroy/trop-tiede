package trop.tiede.lwm2mserver

import org.eclipse.californium.scandium.config.DtlsConnectorConfig
import org.eclipse.leshan.core.californium.PrincipalMdcConnectionListener
import org.eclipse.leshan.core.observation.CompositeObservation
import org.eclipse.leshan.core.observation.Observation
import org.eclipse.leshan.core.observation.SingleObservation
import org.eclipse.leshan.core.response.ObserveCompositeResponse
import org.eclipse.leshan.core.response.ObserveResponse
import org.eclipse.leshan.server.LeshanServer
import org.eclipse.leshan.server.LeshanServerBuilder
import org.eclipse.leshan.server.californium.endpoint.CaliforniumServerEndpointsProvider
import org.eclipse.leshan.server.californium.endpoint.coap.CoapServerProtocolProvider
import org.eclipse.leshan.server.californium.endpoint.coaps.CoapsServerProtocolProvider
import org.eclipse.leshan.server.observation.ObservationListener
import org.eclipse.leshan.server.registration.Registration
import org.eclipse.leshan.server.registration.RegistrationListener
import org.eclipse.leshan.server.registration.RegistrationUpdate
import org.springframework.boot.CommandLineRunner
import org.springframework.boot.SpringApplication
import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.runApplication
import org.springframework.context.ApplicationContext
import org.springframework.context.annotation.Bean
import java.lang.Exception
import java.util.*


@SpringBootApplication
class Lwm2mServerApplication {
	lateinit var server: LeshanServer

	fun main(args: Array<String>) {
		SpringApplication.run(Lwm2mServerApplication::class.java, *args)
	}

	@Bean
	fun commandLineRunner(ctx: ApplicationContext): CommandLineRunner? {
		return CommandLineRunner {
			this.init()
			this.server.start()
		}
	}
	fun init() {
		val builder = LeshanServerBuilder()
		val endpointsBuilder = CaliforniumServerEndpointsProvider.Builder( // Add coap Protocol support
			CoapServerProtocolProvider())  // Add coaps protocol support
		builder.setEndpointsProvider(endpointsBuilder.build())
		this.server = builder.build()
		this.server.observationService.addListener(object : ObservationListener {
			override fun newObservation(observation: Observation?, registration: Registration?) {
				println("new obs $observation.id")
			}

			override fun cancelled(observation: Observation?) {
				TODO("Not yet implemented")
			}

			override fun onResponse(
				observation: SingleObservation?,
				registration: Registration?,
				response: ObserveResponse?
			) {
				TODO("Not yet implemented")
			}

			override fun onResponse(
				observation: CompositeObservation?,
				registration: Registration?,
				response: ObserveCompositeResponse?
			) {
				TODO("Not yet implemented")
			}

			override fun onError(observation: Observation?, registration: Registration?, error: Exception?) {
				TODO("Not yet implemented")
			}

		})
		this.server.registrationService.addListener(object : RegistrationListener {
			override fun registered(
				registration: Registration, previousReg: Registration?,
				previousObsersations: Collection<Observation?>?
			) {
				println("new device: " + registration.endpoint)
			}

			override fun updated(update: RegistrationUpdate?, updatedReg: Registration, previousReg: Registration?) {
				println("device is still here: " + updatedReg.endpoint)
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
