package com.github.blackjack200.xyron;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.56.1)",
    comments = "Source: xchange.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class AnticheatGrpc {

  private AnticheatGrpc() {}

  public static final String SERVICE_NAME = "xchange.Anticheat";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.github.blackjack200.xyron.Xchange.AddPlayerRequest,
      com.github.blackjack200.xyron.Xchange.PlayerReceipt> getAddPlayerMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "AddPlayer",
      requestType = com.github.blackjack200.xyron.Xchange.AddPlayerRequest.class,
      responseType = com.github.blackjack200.xyron.Xchange.PlayerReceipt.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.github.blackjack200.xyron.Xchange.AddPlayerRequest,
      com.github.blackjack200.xyron.Xchange.PlayerReceipt> getAddPlayerMethod() {
    io.grpc.MethodDescriptor<com.github.blackjack200.xyron.Xchange.AddPlayerRequest, com.github.blackjack200.xyron.Xchange.PlayerReceipt> getAddPlayerMethod;
    if ((getAddPlayerMethod = AnticheatGrpc.getAddPlayerMethod) == null) {
      synchronized (AnticheatGrpc.class) {
        if ((getAddPlayerMethod = AnticheatGrpc.getAddPlayerMethod) == null) {
          AnticheatGrpc.getAddPlayerMethod = getAddPlayerMethod =
              io.grpc.MethodDescriptor.<com.github.blackjack200.xyron.Xchange.AddPlayerRequest, com.github.blackjack200.xyron.Xchange.PlayerReceipt>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "AddPlayer"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.github.blackjack200.xyron.Xchange.AddPlayerRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.github.blackjack200.xyron.Xchange.PlayerReceipt.getDefaultInstance()))
              .setSchemaDescriptor(new AnticheatMethodDescriptorSupplier("AddPlayer"))
              .build();
        }
      }
    }
    return getAddPlayerMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.github.blackjack200.xyron.Xchange.PlayerReceipt,
      com.google.protobuf.Empty> getRemovePlayerMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "RemovePlayer",
      requestType = com.github.blackjack200.xyron.Xchange.PlayerReceipt.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.github.blackjack200.xyron.Xchange.PlayerReceipt,
      com.google.protobuf.Empty> getRemovePlayerMethod() {
    io.grpc.MethodDescriptor<com.github.blackjack200.xyron.Xchange.PlayerReceipt, com.google.protobuf.Empty> getRemovePlayerMethod;
    if ((getRemovePlayerMethod = AnticheatGrpc.getRemovePlayerMethod) == null) {
      synchronized (AnticheatGrpc.class) {
        if ((getRemovePlayerMethod = AnticheatGrpc.getRemovePlayerMethod) == null) {
          AnticheatGrpc.getRemovePlayerMethod = getRemovePlayerMethod =
              io.grpc.MethodDescriptor.<com.github.blackjack200.xyron.Xchange.PlayerReceipt, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "RemovePlayer"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.github.blackjack200.xyron.Xchange.PlayerReceipt.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new AnticheatMethodDescriptorSupplier("RemovePlayer"))
              .build();
        }
      }
    }
    return getRemovePlayerMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.github.blackjack200.xyron.Xchange.ReportData,
      com.github.blackjack200.xyron.Xchange.ReportResponse> getReportMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Report",
      requestType = com.github.blackjack200.xyron.Xchange.ReportData.class,
      responseType = com.github.blackjack200.xyron.Xchange.ReportResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.github.blackjack200.xyron.Xchange.ReportData,
      com.github.blackjack200.xyron.Xchange.ReportResponse> getReportMethod() {
    io.grpc.MethodDescriptor<com.github.blackjack200.xyron.Xchange.ReportData, com.github.blackjack200.xyron.Xchange.ReportResponse> getReportMethod;
    if ((getReportMethod = AnticheatGrpc.getReportMethod) == null) {
      synchronized (AnticheatGrpc.class) {
        if ((getReportMethod = AnticheatGrpc.getReportMethod) == null) {
          AnticheatGrpc.getReportMethod = getReportMethod =
              io.grpc.MethodDescriptor.<com.github.blackjack200.xyron.Xchange.ReportData, com.github.blackjack200.xyron.Xchange.ReportResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Report"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.github.blackjack200.xyron.Xchange.ReportData.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.github.blackjack200.xyron.Xchange.ReportResponse.getDefaultInstance()))
              .setSchemaDescriptor(new AnticheatMethodDescriptorSupplier("Report"))
              .build();
        }
      }
    }
    return getReportMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.github.blackjack200.xyron.Xchange.BatchedReportData,
      com.github.blackjack200.xyron.Xchange.BatchedReportResponse> getReportBatchedMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ReportBatched",
      requestType = com.github.blackjack200.xyron.Xchange.BatchedReportData.class,
      responseType = com.github.blackjack200.xyron.Xchange.BatchedReportResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.github.blackjack200.xyron.Xchange.BatchedReportData,
      com.github.blackjack200.xyron.Xchange.BatchedReportResponse> getReportBatchedMethod() {
    io.grpc.MethodDescriptor<com.github.blackjack200.xyron.Xchange.BatchedReportData, com.github.blackjack200.xyron.Xchange.BatchedReportResponse> getReportBatchedMethod;
    if ((getReportBatchedMethod = AnticheatGrpc.getReportBatchedMethod) == null) {
      synchronized (AnticheatGrpc.class) {
        if ((getReportBatchedMethod = AnticheatGrpc.getReportBatchedMethod) == null) {
          AnticheatGrpc.getReportBatchedMethod = getReportBatchedMethod =
              io.grpc.MethodDescriptor.<com.github.blackjack200.xyron.Xchange.BatchedReportData, com.github.blackjack200.xyron.Xchange.BatchedReportResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ReportBatched"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.github.blackjack200.xyron.Xchange.BatchedReportData.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.github.blackjack200.xyron.Xchange.BatchedReportResponse.getDefaultInstance()))
              .setSchemaDescriptor(new AnticheatMethodDescriptorSupplier("ReportBatched"))
              .build();
        }
      }
    }
    return getReportBatchedMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static AnticheatStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<AnticheatStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<AnticheatStub>() {
        @java.lang.Override
        public AnticheatStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new AnticheatStub(channel, callOptions);
        }
      };
    return AnticheatStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static AnticheatBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<AnticheatBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<AnticheatBlockingStub>() {
        @java.lang.Override
        public AnticheatBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new AnticheatBlockingStub(channel, callOptions);
        }
      };
    return AnticheatBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static AnticheatFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<AnticheatFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<AnticheatFutureStub>() {
        @java.lang.Override
        public AnticheatFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new AnticheatFutureStub(channel, callOptions);
        }
      };
    return AnticheatFutureStub.newStub(factory, channel);
  }

  /**
   */
  public interface AsyncService {

    /**
     */
    default void addPlayer(com.github.blackjack200.xyron.Xchange.AddPlayerRequest request,
        io.grpc.stub.StreamObserver<com.github.blackjack200.xyron.Xchange.PlayerReceipt> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getAddPlayerMethod(), responseObserver);
    }

    /**
     */
    default void removePlayer(com.github.blackjack200.xyron.Xchange.PlayerReceipt request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getRemovePlayerMethod(), responseObserver);
    }

    /**
     */
    default void report(com.github.blackjack200.xyron.Xchange.ReportData request,
        io.grpc.stub.StreamObserver<com.github.blackjack200.xyron.Xchange.ReportResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getReportMethod(), responseObserver);
    }

    /**
     */
    default void reportBatched(com.github.blackjack200.xyron.Xchange.BatchedReportData request,
        io.grpc.stub.StreamObserver<com.github.blackjack200.xyron.Xchange.BatchedReportResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getReportBatchedMethod(), responseObserver);
    }
  }

  /**
   * Base class for the server implementation of the service Anticheat.
   */
  public static abstract class AnticheatImplBase
      implements io.grpc.BindableService, AsyncService {

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return AnticheatGrpc.bindService(this);
    }
  }

  /**
   * A stub to allow clients to do asynchronous rpc calls to service Anticheat.
   */
  public static final class AnticheatStub
      extends io.grpc.stub.AbstractAsyncStub<AnticheatStub> {
    private AnticheatStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected AnticheatStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new AnticheatStub(channel, callOptions);
    }

    /**
     */
    public void addPlayer(com.github.blackjack200.xyron.Xchange.AddPlayerRequest request,
        io.grpc.stub.StreamObserver<com.github.blackjack200.xyron.Xchange.PlayerReceipt> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getAddPlayerMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void removePlayer(com.github.blackjack200.xyron.Xchange.PlayerReceipt request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getRemovePlayerMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void report(com.github.blackjack200.xyron.Xchange.ReportData request,
        io.grpc.stub.StreamObserver<com.github.blackjack200.xyron.Xchange.ReportResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getReportMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void reportBatched(com.github.blackjack200.xyron.Xchange.BatchedReportData request,
        io.grpc.stub.StreamObserver<com.github.blackjack200.xyron.Xchange.BatchedReportResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getReportBatchedMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * A stub to allow clients to do synchronous rpc calls to service Anticheat.
   */
  public static final class AnticheatBlockingStub
      extends io.grpc.stub.AbstractBlockingStub<AnticheatBlockingStub> {
    private AnticheatBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected AnticheatBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new AnticheatBlockingStub(channel, callOptions);
    }

    /**
     */
    public com.github.blackjack200.xyron.Xchange.PlayerReceipt addPlayer(com.github.blackjack200.xyron.Xchange.AddPlayerRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getAddPlayerMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.google.protobuf.Empty removePlayer(com.github.blackjack200.xyron.Xchange.PlayerReceipt request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getRemovePlayerMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.github.blackjack200.xyron.Xchange.ReportResponse report(com.github.blackjack200.xyron.Xchange.ReportData request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getReportMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.github.blackjack200.xyron.Xchange.BatchedReportResponse reportBatched(com.github.blackjack200.xyron.Xchange.BatchedReportData request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getReportBatchedMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do ListenableFuture-style rpc calls to service Anticheat.
   */
  public static final class AnticheatFutureStub
      extends io.grpc.stub.AbstractFutureStub<AnticheatFutureStub> {
    private AnticheatFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected AnticheatFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new AnticheatFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.github.blackjack200.xyron.Xchange.PlayerReceipt> addPlayer(
        com.github.blackjack200.xyron.Xchange.AddPlayerRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getAddPlayerMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> removePlayer(
        com.github.blackjack200.xyron.Xchange.PlayerReceipt request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getRemovePlayerMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.github.blackjack200.xyron.Xchange.ReportResponse> report(
        com.github.blackjack200.xyron.Xchange.ReportData request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getReportMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.github.blackjack200.xyron.Xchange.BatchedReportResponse> reportBatched(
        com.github.blackjack200.xyron.Xchange.BatchedReportData request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getReportBatchedMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_ADD_PLAYER = 0;
  private static final int METHODID_REMOVE_PLAYER = 1;
  private static final int METHODID_REPORT = 2;
  private static final int METHODID_REPORT_BATCHED = 3;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final AsyncService serviceImpl;
    private final int methodId;

    MethodHandlers(AsyncService serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_ADD_PLAYER:
          serviceImpl.addPlayer((com.github.blackjack200.xyron.Xchange.AddPlayerRequest) request,
              (io.grpc.stub.StreamObserver<com.github.blackjack200.xyron.Xchange.PlayerReceipt>) responseObserver);
          break;
        case METHODID_REMOVE_PLAYER:
          serviceImpl.removePlayer((com.github.blackjack200.xyron.Xchange.PlayerReceipt) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_REPORT:
          serviceImpl.report((com.github.blackjack200.xyron.Xchange.ReportData) request,
              (io.grpc.stub.StreamObserver<com.github.blackjack200.xyron.Xchange.ReportResponse>) responseObserver);
          break;
        case METHODID_REPORT_BATCHED:
          serviceImpl.reportBatched((com.github.blackjack200.xyron.Xchange.BatchedReportData) request,
              (io.grpc.stub.StreamObserver<com.github.blackjack200.xyron.Xchange.BatchedReportResponse>) responseObserver);
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

  public static final io.grpc.ServerServiceDefinition bindService(AsyncService service) {
    return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
        .addMethod(
          getAddPlayerMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.github.blackjack200.xyron.Xchange.AddPlayerRequest,
              com.github.blackjack200.xyron.Xchange.PlayerReceipt>(
                service, METHODID_ADD_PLAYER)))
        .addMethod(
          getRemovePlayerMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.github.blackjack200.xyron.Xchange.PlayerReceipt,
              com.google.protobuf.Empty>(
                service, METHODID_REMOVE_PLAYER)))
        .addMethod(
          getReportMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.github.blackjack200.xyron.Xchange.ReportData,
              com.github.blackjack200.xyron.Xchange.ReportResponse>(
                service, METHODID_REPORT)))
        .addMethod(
          getReportBatchedMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.github.blackjack200.xyron.Xchange.BatchedReportData,
              com.github.blackjack200.xyron.Xchange.BatchedReportResponse>(
                service, METHODID_REPORT_BATCHED)))
        .build();
  }

  private static abstract class AnticheatBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    AnticheatBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return com.github.blackjack200.xyron.Xchange.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("Anticheat");
    }
  }

  private static final class AnticheatFileDescriptorSupplier
      extends AnticheatBaseDescriptorSupplier {
    AnticheatFileDescriptorSupplier() {}
  }

  private static final class AnticheatMethodDescriptorSupplier
      extends AnticheatBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    AnticheatMethodDescriptorSupplier(String methodName) {
      this.methodName = methodName;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.MethodDescriptor getMethodDescriptor() {
      return getServiceDescriptor().findMethodByName(methodName);
    }
  }

  private static volatile io.grpc.ServiceDescriptor serviceDescriptor;

  public static io.grpc.ServiceDescriptor getServiceDescriptor() {
    io.grpc.ServiceDescriptor result = serviceDescriptor;
    if (result == null) {
      synchronized (AnticheatGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new AnticheatFileDescriptorSupplier())
              .addMethod(getAddPlayerMethod())
              .addMethod(getRemovePlayerMethod())
              .addMethod(getReportMethod())
              .addMethod(getReportBatchedMethod())
              .build();
        }
      }
    }
    return result;
  }
}
