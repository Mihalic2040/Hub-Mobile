package com.Bridge.protocols;

import static io.grpc.MethodDescriptor.generateFullMethodName;
import static io.grpc.stub.ClientCalls.asyncBidiStreamingCall;
import static io.grpc.stub.ClientCalls.asyncClientStreamingCall;
import static io.grpc.stub.ClientCalls.asyncServerStreamingCall;
import static io.grpc.stub.ClientCalls.asyncUnaryCall;
import static io.grpc.stub.ClientCalls.blockingServerStreamingCall;
import static io.grpc.stub.ClientCalls.blockingUnaryCall;
import static io.grpc.stub.ClientCalls.futureUnaryCall;
import static io.grpc.stub.ServerCalls.asyncBidiStreamingCall;
import static io.grpc.stub.ServerCalls.asyncClientStreamingCall;
import static io.grpc.stub.ServerCalls.asyncServerStreamingCall;
import static io.grpc.stub.ServerCalls.asyncUnaryCall;
import static io.grpc.stub.ServerCalls.asyncUnimplementedStreamingCall;
import static io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.28.1)",
    comments = "Source: protocols/bridge.proto")
public final class PeersGrpc {

  private PeersGrpc() {}

  public static final String SERVICE_NAME = "Peers";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.Bridge.protocols.PeersResponse,
      com.Bridge.protocols.PeersResponse> getSayHelloMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "SayHello",
      requestType = com.Bridge.protocols.PeersResponse.class,
      responseType = com.Bridge.protocols.PeersResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.Bridge.protocols.PeersResponse,
      com.Bridge.protocols.PeersResponse> getSayHelloMethod() {
    io.grpc.MethodDescriptor<com.Bridge.protocols.PeersResponse, com.Bridge.protocols.PeersResponse> getSayHelloMethod;
    if ((getSayHelloMethod = PeersGrpc.getSayHelloMethod) == null) {
      synchronized (PeersGrpc.class) {
        if ((getSayHelloMethod = PeersGrpc.getSayHelloMethod) == null) {
          PeersGrpc.getSayHelloMethod = getSayHelloMethod =
              io.grpc.MethodDescriptor.<com.Bridge.protocols.PeersResponse, com.Bridge.protocols.PeersResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "SayHello"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  com.Bridge.protocols.PeersResponse.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  com.Bridge.protocols.PeersResponse.getDefaultInstance()))
              .build();
        }
      }
    }
    return getSayHelloMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static PeersStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<PeersStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<PeersStub>() {
        @java.lang.Override
        public PeersStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new PeersStub(channel, callOptions);
        }
      };
    return PeersStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static PeersBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<PeersBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<PeersBlockingStub>() {
        @java.lang.Override
        public PeersBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new PeersBlockingStub(channel, callOptions);
        }
      };
    return PeersBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static PeersFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<PeersFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<PeersFutureStub>() {
        @java.lang.Override
        public PeersFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new PeersFutureStub(channel, callOptions);
        }
      };
    return PeersFutureStub.newStub(factory, channel);
  }

  /**
   */
  public static abstract class PeersImplBase implements io.grpc.BindableService {

    /**
     */
    public void sayHello(com.Bridge.protocols.PeersResponse request,
        io.grpc.stub.StreamObserver<com.Bridge.protocols.PeersResponse> responseObserver) {
      asyncUnimplementedUnaryCall(getSayHelloMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getSayHelloMethod(),
            asyncUnaryCall(
              new MethodHandlers<
                com.Bridge.protocols.PeersResponse,
                com.Bridge.protocols.PeersResponse>(
                  this, METHODID_SAY_HELLO)))
          .build();
    }
  }

  /**
   */
  public static final class PeersStub extends io.grpc.stub.AbstractAsyncStub<PeersStub> {
    private PeersStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected PeersStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new PeersStub(channel, callOptions);
    }

    /**
     */
    public void sayHello(com.Bridge.protocols.PeersResponse request,
        io.grpc.stub.StreamObserver<com.Bridge.protocols.PeersResponse> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(getSayHelloMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class PeersBlockingStub extends io.grpc.stub.AbstractBlockingStub<PeersBlockingStub> {
    private PeersBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected PeersBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new PeersBlockingStub(channel, callOptions);
    }

    /**
     */
    public com.Bridge.protocols.PeersResponse sayHello(com.Bridge.protocols.PeersResponse request) {
      return blockingUnaryCall(
          getChannel(), getSayHelloMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class PeersFutureStub extends io.grpc.stub.AbstractFutureStub<PeersFutureStub> {
    private PeersFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected PeersFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new PeersFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.Bridge.protocols.PeersResponse> sayHello(
        com.Bridge.protocols.PeersResponse request) {
      return futureUnaryCall(
          getChannel().newCall(getSayHelloMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_SAY_HELLO = 0;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final PeersImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(PeersImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_SAY_HELLO:
          serviceImpl.sayHello((com.Bridge.protocols.PeersResponse) request,
              (io.grpc.stub.StreamObserver<com.Bridge.protocols.PeersResponse>) responseObserver);
          break;
        default:
          throw new AssertionError();
      }
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public io.grpc.stub.StreamObserver<Req> invoke(
        io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        default:
          throw new AssertionError();
      }
    }
  }

  private static volatile io.grpc.ServiceDescriptor serviceDescriptor;

  public static io.grpc.ServiceDescriptor getServiceDescriptor() {
    io.grpc.ServiceDescriptor result = serviceDescriptor;
    if (result == null) {
      synchronized (PeersGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .addMethod(getSayHelloMethod())
              .build();
        }
      }
    }
    return result;
  }
}
