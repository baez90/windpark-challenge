using Microsoft.EntityFrameworkCore;
using WindparkChallenge.API.Consumer;

namespace WindparkChallenge.API.DB;

public class StatsContext : DbContext
{
    public StatsContext(DbContextOptions options) : base(options)
    {
    }

    public DbSet<Site> Sites { get; set; }

    public DbSet<Turbine> Turbines { get; set; }

    public DbSet<Stats> Stats { get; set; }

    protected override void OnModelCreating(ModelBuilder modelBuilder)
    {
        modelBuilder.Entity<Site>()
            .HasMany(p => p.Turbines)
            .WithOne(t => t.Site);
        
        modelBuilder.Entity<Turbine>()
            .HasMany(t => t.TurbineStats)
            .WithOne(s => s.Turbine);
    }
}