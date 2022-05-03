// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

package tailerfactory

/*

// reset clears the factories slice and resets the sync.Once.
func reset() {
	factories = factorySlice{}
	sortFactoriesOnce = sync.Once{}
}

type fakeTailer struct {
	name string
}

// Start implements tailers.Tailer#Start.
func (*fakeTailer) Start(pipeline chan *message.Message, registry auditor.Registry) {}

// Stop implements tailers.Tailer#Stop.
func (*fakeTailer) Stop() {}

func TestFactoryMatches(t *testing.T) {
	reset()
	called := false
	registerFactory(func(src *sources.LogSource) tailers.Tailer {
		called = true
		require.Equal(t, src.Name, "test")
		return &fakeTailer{name: "a"}
	}, 1)

	tailer := MakeTailerForSource(&sources.LogSource{Name: "test"})
	require.True(t, called)
	require.NotNil(t, tailer)
	require.Equal(t, "a", (tailer.(*fakeTailer)).name)
}

func TestPriority(t *testing.T) {
	reset()
	registerFactory(func(src *sources.LogSource) tailers.Tailer { panic("wrong factory 2") }, 2)
	registerFactory(func(src *sources.LogSource) tailers.Tailer { return &fakeTailer{name: "a"} }, 1)
	registerFactory(func(src *sources.LogSource) tailers.Tailer { panic("wrong factory 3") }, 3)

	tailer := MakeTailerForSource(&sources.LogSource{})
	require.NotNil(t, tailer)
	require.Equal(t, "a", (tailer.(*fakeTailer)).name)
}

func TestFallback(t *testing.T) {
	reset()
	calls := []int{}
	registerFactory(func(src *sources.LogSource) tailers.Tailer { calls = append(calls, 1); return nil }, 1)
	registerFactory(func(src *sources.LogSource) tailers.Tailer { return &fakeTailer{name: "4"} }, 4)
	registerFactory(func(src *sources.LogSource) tailers.Tailer { calls = append(calls, 2); return nil }, 2)
	registerFactory(func(src *sources.LogSource) tailers.Tailer { calls = append(calls, 5); return nil }, 5)
	registerFactory(func(src *sources.LogSource) tailers.Tailer { calls = append(calls, 3); return nil }, 3)

	tailer := MakeTailerForSource(&sources.LogSource{})
	require.NotNil(t, tailer)
	require.Equal(t, []int{1, 2, 3}, calls)
	require.Equal(t, "4", (tailer.(*fakeTailer)).name)
}

func TestFallbackNone(t *testing.T) {
	reset()
	calls := []int{}
	registerFactory(func(src *sources.LogSource) tailers.Tailer { calls = append(calls, 4); return nil }, 4)
	registerFactory(func(src *sources.LogSource) tailers.Tailer { calls = append(calls, 3); return nil }, 3)
	registerFactory(func(src *sources.LogSource) tailers.Tailer { calls = append(calls, 5); return nil }, 5)
	registerFactory(func(src *sources.LogSource) tailers.Tailer { calls = append(calls, 1); return nil }, 1)
	registerFactory(func(src *sources.LogSource) tailers.Tailer { calls = append(calls, 2); return nil }, 2)

	tailer := MakeTailerForSource(&sources.LogSource{})
	require.Nil(t, tailer)
	require.Equal(t, []int{1, 2, 3, 4, 5}, calls)
}
*/
