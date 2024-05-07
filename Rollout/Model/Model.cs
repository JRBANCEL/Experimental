namespace Model;

public sealed class SemVer
{
    public int Major { get; init; }
    public int Minor { get; init; }
    public int Patch { get; init; }
    public string? PreRelease { get; init; }
    public string? BuildMetadata { get; init; }
    
    public SemVer(string version)
    {
        var parts = version.Split('.');
        Major = int.Parse(parts[0]);
        Minor = int.Parse(parts[1]);
        Patch = int.Parse(parts[2]);

        if (parts.Length <= 3) return;

        var preReleaseParts = parts[3].Split('-');
        PreRelease = preReleaseParts[0];
        if (preReleaseParts.Length > 1)
        {
            BuildMetadata = preReleaseParts[1];
        }
    }
    
    public override string ToString()
    {
        return $"{Major}.{Minor}.{Patch}{(PreRelease != null ? "-" + PreRelease : "")}{(BuildMetadata != null ? "+" + BuildMetadata : "")}";
    }
}

public sealed class Release
{
    public SemVer Version { get; init; }
    
    public IDictionary<string, string> Metadata { get; init; }
}


public interface IReleaseLister
{
    Task<IEnumerable<Release>> ListReleasesAsync(CancellationToken cancellationToken = default);
}

public sealed class Target
{
    public string Name { get; init; }
    
    public IDictionary<string, string> Metadata { get; init; }
}

public interface ITargetLister
{
    Task<IEnumerable<Target>> ListTargetsAsync(CancellationToken cancellationToken = default);
}

public sealed class TargetStatus
{
//    public State State { get; init; }
    
    public SemVer Version { get; init; }
    
//    public DateTime? LastUpdated { get; init; }
}

public sealed class PlanStep
{
    public string MetadataKey { get; init; }
    
    public IList<string>? Order { get; init; }
}

public sealed class Plan
{
    public IList<PlanStep> Steps { get; init; }
}

public sealed class Project
{
    public IReleaseLister ReleaseLister { get; init; }
    
    public ITargetLister TargetLister { get; init; }
    
    public IList<Plan> Plans { get; init; }
    
    public async Task RenderAsync(CancellationToken cancellationToken = default)
    {
        var releases = await this.ReleaseLister.ListReleasesAsync(cancellationToken);
        var targets = await this.TargetLister.ListTargetsAsync(cancellationToken);
        
        foreach (var target in targets)
        {
            foreach (var plan in Plans)
            {
                foreach (var step in plan.Steps)
                {
                    if (step.Order != null)
                    {
                        foreach (var order in step.Order)
                        {
                            if (target.Metadata.TryGetValue(step.MetadataKey, out var value) && value == order)
                            {
                                // Render
                            }
                        }
                    }
                    else
                    {
                        if (target.Metadata.TryGetValue(step.MetadataKey, out var value))
                        {
                            // Render
                        }
                    }
                }
            }
        }
    }
}

public interface IConstraint
{
    Task<bool> IsSatisfiedAsync(Target target, TargetStatus status, CancellationToken cancellationToken = default);
}
