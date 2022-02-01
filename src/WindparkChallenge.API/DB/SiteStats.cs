namespace WindparkChallenge.API.DB;

public class Site
{
    public int Id { get; set; }

    public string Name { get; set; }

    public ICollection<Turbine> Turbines { get; set; }
}

public class Turbine
{
    public int Id { get; set; }

    public string Name { get; set; }

    public Site Site { get; set; }

    public ICollection<Stats> TurbineStats { get; set; }
}

public class Stats
{
    public int Id { get; set; }

    public Turbine Turbine { get; set; }

    public DateTimeOffset Timestamp { get; set; }

    public double WindSpeed { get; set; }

    public double CurrentProduction { get; set; }
}