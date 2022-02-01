using System.Text.Json;
using RabbitMQ.Client;
using RabbitMQ.Client.Events;

namespace WindparkChallenge.API.Consumer;

public class RabbitMQConsumer : IHostedService
{
    private readonly ILogger<RabbitMQConsumer> _logger;
    private IConnection? _connection;
    private IModel? _channel;

    public RabbitMQConsumer(ILogger<RabbitMQConsumer> logger)
    {
        _logger = logger;
    }

    public Task StartAsync(CancellationToken cancellationToken)
    {
        var connectionFactory = new ConnectionFactory
        {
            Uri = new Uri("amqp://rabbitmq:rabbitmq@localhost/"),
            DispatchConsumersAsync = true
        };

        _connection = connectionFactory.CreateConnection();
        _channel = _connection.CreateModel();

        var consumer = new AsyncEventingBasicConsumer(_channel);
        consumer.Received += ConsumeMessagesAsync;
        _channel.BasicConsume("windpark-stats", true, consumer);
        
        return Task.CompletedTask;
    }

    public Task StopAsync(CancellationToken cancellationToken)
    {
        _connection?.Close();
        _connection?.Dispose();
        return Task.CompletedTask;
    }

    private async Task ConsumeMessagesAsync(object ch, BasicDeliverEventArgs args)
    {
        var snapshot = JsonSerializer.Deserialize<ParksSnapshot>(args.Body.Span);
        _logger.LogInformation("Got snapshot: {Snapshot}", snapshot);
    }
}