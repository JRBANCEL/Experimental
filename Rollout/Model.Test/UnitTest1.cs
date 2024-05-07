using System.Globalization;
using Microsoft.VisualBasic.CompilerServices;

namespace Model.Test;

[TestClass]
public class UnitTest1
{
    [TestMethod]
    public void TestModel()
    {
        var project = new Project
        {
            ReleaseLister = new StaticReleaseLister(),
            TargetLister = new StaticTargetLister(),
            Plans = new List<Plan>
            {
                new()
                {
                    Steps = new List<PlanStep>
                    {
                        new()
                        {
                            MetadataKey = "Environment",
                            Order = new List<string>
                            {
                                "Alpha",
                                "Staging",
                                "Production"
                            }
                        },
                        new()
                        {
                            MetadataKey = "Region"
                        },
                        new()
                        {
                            MetadataKey = "ScaleUnit"
                        }
                    }
                }
            }
        };
        
        
    }
}

public class SequentialVersionConstraint : IConstraint
{
    public Task<bool> IsSatisfiedAsync(Target target, TargetStatus status, CancellationToken cancellationToken = default)
    {
        // Needs access to new version
        throw new NotImplementedException();
    }
}

public class StaticReleaseLister : IReleaseLister
{
    public Task<IEnumerable<Release>> ListReleasesAsync(CancellationToken cancellationToken = default)
    {
        var releases = new []
        {
            new Release
            {
                Version = new SemVer("1.0.0")
            },
            new Release
            {
                Version = new SemVer("1.0.1")
            },
            new Release
            {
                Version = new SemVer("1.1.0")
            },
        };

        return Task.FromResult(releases.AsEnumerable());
    }
}

public class StaticTargetLister : ITargetLister
{
    public Task<IEnumerable<Target>> ListTargetsAsync(CancellationToken cancellationToken = default)
    {
        var targets = new []
        {
            new Target
            {
                Name = "alp-eastus-0",
                Metadata = new Dictionary<string, string>
                {
                    ["Environment"] = "Alpha",
                    ["Region"] = "eastus-0",
                }
            },
            new Target
            {
                Name = "alp-westus-2",
                Metadata = new Dictionary<string, string>
                {
                    ["Environment"] = "Alpha",
                    ["Region"] = "westus-2",
                }
            },
            new Target
            {
                Name = "prd-eastus-0",
                Metadata = new Dictionary<string, string>
                {
                    ["Environment"] = "Production",
                    ["Region"] = "eastus-0",
                }
            },
        };

        return Task.FromResult(targets.AsEnumerable());
    }
}
