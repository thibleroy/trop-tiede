import org.eclipse.leshan.core.observation.Observation
import org.eclipse.leshan.server.californium.LeshanServer
import org.eclipse.leshan.server.californium.LeshanServerBuilder
import org.eclipse.leshan.server.registration.Registration
import org.eclipse.leshan.server.registration.RegistrationListener
import org.eclipse.leshan.server.registration.RegistrationUpdate


fun main(args: Array<String>) {
    println("Hello World! 3")
    val builder = LeshanServerBuilder()
    val server: LeshanServer = builder.build()
    server.registrationService.addListener(object : RegistrationListener {
        override fun registered(
            registration: Registration, previousReg: Registration?,
            previousObsersations: Collection<Observation?>?
        ) {
            println("new device: " + registration.endpoint)
        }

        override fun updated(update: RegistrationUpdate, updatedReg: Registration, previousReg: Registration) {
            println("device is still here: " + updatedReg.endpoint)
        }

        override fun unregistered(
            registration: Registration, observations: Collection<Observation?>?, expired: Boolean,
            newReg: Registration?
        ) {
            println("device left: " + registration.endpoint)
        }
    })
    server.start()
}