namespace WindparkChallenge.API.Consumer;

public record StatValue
{
    public double WindSpeed { get; set; }

    public double CurrentProduction { get; set; }
}

public record Turbine
{
    public long Id { get; set; }

    public string Name { get; set; }
    public StatValue Stats { get; set; }
}

public record Park
{
    public long Id { get; set; }

    public string Name { get; set; }
    public IEnumerable<Turbine>? Turbines { get; set; }
}

public record ParksSnapshot
{
    public DateTimeOffset Timestamp { get; init; }

    public IEnumerable<Park>? Parks { get; init; }
}