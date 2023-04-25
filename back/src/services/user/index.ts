import 'dotenv/config';
import { Connection, ConsumeMessage, connect, Options } from 'amqplib';
import { generateUuid, publishRPCResponse, serveRPC } from './rpc.js';

const options: Options.Connect = {
  hostname: process.env.RABBITMQ_BROKER_URL,
  port: parseInt(process.env.RABBITMQ_BROKER_PORT!),
  username: process.env.RABBITMQ_BROKER_USERNAME,
  password: process.env.RABBITMQ_BROKER_PASSWORD
}
console.log("opts", options);
interface User {
  id: string;
  firstname: string;
}
const connection: Connection = await connect(
  options
)

const handleRPC = async(msg: ConsumeMessage): Promise<void> => {
  const user: User = {
    id: generateUuid(),
    firstname: "test"
}
await publishRPCResponse(connection, msg, JSON.stringify(user));

}

serveRPC(connection, 'user_rpc', handleRPC)
