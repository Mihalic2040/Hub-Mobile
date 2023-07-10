package com.Bridge.api;



import com.Bridge.protocols.RequestData;
import com.Bridge.protocols.ResponseData;
import com.Bridge.protocols.RequestServiceGrpc;

import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import io.grpc.StatusRuntimeException;

public class Request {

    public ResponseData request(RequestData r) {
        // Set up a connection to the gRPC server
        ManagedChannel channel = ManagedChannelBuilder.forAddress(settings.host, settings.port)
                .usePlaintext()
                .build();
        try {
            // Create a gRPC client
            RequestServiceGrpc.RequestServiceBlockingStub client = RequestServiceGrpc.newBlockingStub(channel);


            // Prepare the request

            // Send the gRPC request
            ResponseData response = client.request(r);

            // Process the response
            return response;
        } catch (StatusRuntimeException e) {
            System.err.println("RPC failed: " + e.getStatus());
            return null;
        } finally {
            // Close the channel
            channel.shutdownNow();
        }
    }

}
