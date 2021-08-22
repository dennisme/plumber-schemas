// package: protos.backends
// file: backends/azure-service-bus.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";

export class AzureServiceBus extends jspb.Message { 
    getQueue(): string;
    setQueue(value: string): AzureServiceBus;
    getTopic(): string;
    setTopic(value: string): AzureServiceBus;
    getSubscriptionName(): string;
    setSubscriptionName(value: string): AzureServiceBus;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): AzureServiceBus.AsObject;
    static toObject(includeInstance: boolean, msg: AzureServiceBus): AzureServiceBus.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: AzureServiceBus, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): AzureServiceBus;
    static deserializeBinaryFromReader(message: AzureServiceBus, reader: jspb.BinaryReader): AzureServiceBus;
}

export namespace AzureServiceBus {
    export type AsObject = {
        queue: string,
        topic: string,
        subscriptionName: string,
    }
}
