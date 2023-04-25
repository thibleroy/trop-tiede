import { Connection, ConsumeMessage } from 'amqplib'
export const generateUuid = (): string => {
    return Math.random().toString() +
           Math.random().toString() +
           Math.random().toString();
  }

export const RPC = async(connection: Connection, message: string, rpc_queue_name: string, rpc_cb_queue_name: string, correlation_id: string): Promise<any> => {
	const channel = await connection.createChannel()
    if (!!channel) {
        const assertQueue = await channel.assertQueue(rpc_cb_queue_name, {exclusive: true})
        if (!!assertQueue) {
          channel.consume(assertQueue.queue, msg => {
            if (msg!.properties.correlationId === correlation_id) {
              return msg!.content.toString()
            }
          }, {
            noAck: true
          });
          channel.sendToQueue(rpc_queue_name,
            Buffer.from(message),{
              correlationId: correlation_id,
              replyTo: assertQueue.queue });

        }
    }
}

type RPCRequestHandler = (msg: ConsumeMessage) => void;

export const publishRPCResponse = async(connection: Connection, messageReceived: ConsumeMessage, messageToSend: string) => {
const channel = await connection.createChannel()
    if (!!channel) {
      console.log("msg received", messageReceived.content.toString())
          channel.sendToQueue(messageReceived!.properties.replyTo,
            Buffer.from(messageToSend), {
              correlationId: messageReceived!.properties.correlationId
            });
            await channel.close();
        }
  }

export const serveRPC = async(connection: Connection, rpc_queue_name: string, handler: RPCRequestHandler) => {
    const channel = await connection.createChannel()
    if (!!channel) {
        await channel.assertQueue(rpc_queue_name, {
            durable: false
          });
          //channel.prefetch(1);
        console.log(' [x] Awaiting RPC requests');
        await channel.consume(rpc_queue_name, msg => {
            handler(msg!);
          });
    }
}

