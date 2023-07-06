package com.Bridge.api;

import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import io.grpc.StatusRuntimeException;
import com.Bridge.protocols.PeersGrpc;
import com.Bridge.protocols.PeersResponse;

public class Peers {
    public PeersResponse request() {
        // Set up a connection to the gRPC server
        ManagedChannel channel = ManagedChannelBuilder.forAddress(settings.host, settings.port)
                .usePlaintext()
                .build();
        try {
            // Create a gRPC client
            PeersGrpc.PeersBlockingStub client = PeersGrpc.newBlockingStub(channel);

            // Prepare the request
            PeersResponse request = PeersResponse.newBuilder()
                    .setPeers("")
                    .build();

            // Send the gRPC request
            PeersResponse response = client.sayHello(request);

            // Process the response
            System.out.println("Response GRPC: " + response.getPeers());
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