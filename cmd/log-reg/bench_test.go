package main


import (
        "fmt"
        "testing"

        pb "k8s.io/cri-api/pkg/apis/runtime/v1"
)

func buildResponse(n int) *pb.ListContainerStatsResponse {
        stats := make([]*pb.ContainerStats, n)
        for i := range stats {
                stats[i] = &pb.ContainerStats{
                        Attributes: &pb.ContainerAttributes{
                                Id:       fmt.Sprintf("container-%d-abcdef1234567890", i),
                                Metadata: &pb.ContainerMetadata{Name: fmt.Sprintf("nginx-%d", i)},
                        },
                        Cpu: &pb.CpuUsage{
                                UsageCoreNanoSeconds: &pb.UInt64Value{Value: uint64(i * 38789679)},
                                UsageNanoCores:       &pb.UInt64Value{Value: uint64(i * 3572)},
                        },
                        Memory: &pb.MemoryUsage{
                                WorkingSetBytes: &pb.UInt64Value{Value: 12705792},
                                UsageBytes:      &pb.UInt64Value{Value: 13258752},
                        },
                }
        }
        return &pb.ListContainerStatsResponse{Stats: stats}
}

// Before the change: %#v (fast Go formatting)
func BenchmarkBeforeChange_1000(b *testing.B) {
        r := buildResponse(1000)
        for range b.N {
                _ = fmt.Sprintf("Response: %#v", r)
        }
}

// After the change: %+v (slow protobuf text formatting)
func BenchmarkAfterChange_1000(b *testing.B) {
        r := buildResponse(1000)
        for range b.N {
                _ = fmt.Sprintf("Response: %T: %+v", r, r)
        }
}

// After the fix: %#v (fast again)
func BenchmarkAfterFix_1000(b *testing.B) {
        r := buildResponse(1000)
        for range b.N {
                _ = fmt.Sprintf("Response: %#v", r)
        }
}
