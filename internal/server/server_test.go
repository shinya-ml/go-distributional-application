package server

// func TestServer(t *testing.T) {
// 	for scenario, fn := range map[string]func(
// 		t *testing.T,
// 		client api.LogClient,
// 		config *Config,
// 	){
// 		"produce/consume a message to/from the log succeeds": testProduceConsume,
// 		"produce/consume stream succeeds":                    testProduceConsumeStream,
// 		"consume past log boundary fails":                    testConsumePastBoundary,
// 	} {
// 		t.Run(scenario, func(t *testing.T) {
// 			client, config, teardown := setupTest(t, nil)
// 			defer teardown()
// 			fn(t, client, config)
// 		})
// 	}
// }

// func setupTest(t *testing.T, fn func(*Config)) (
// 	client api.LogClient,
// 	cfg *Config,
// 	teardown func(),
// ) {
// 	t.Helper()

// 	l, err := net.Listen("tcp", ":0")
// 	require.NoError(t, err)

// 	clientOptions := []grpc.DialOption{
// 		grpc.WithTransportCredentials(insecure.NewCredentials()),
// 	}
// 	cc, err := grpc.Dial(l.Addr().String(), clientOptions...)
// 	require.NoError(t, err)

// 	dir, err := os.MkdirTemp("", "server-test")
// 	require.NoError(t, err)

// 	clog, err := log.NewLog(dir, log.Config{})
// 	require.NoError(t, err)

// 	cfg = &Config{
// 		CommitLog: clog,
// 	}
// 	if fn != nil {
// 		fn(cfg)
// 	}
// 	server, err := NewGRPCServer(cfg)
// 	require.NoError(t, err)

// 	go func() {
// 		server.Serve(l)
// 	}()

// 	client = api.NewLogClient(cc)
// 	return client, cfg, func() {
// 		cc.Close()
// 		server.Stop()
// 		l.Close()
// 		clog.Remove()
// 	}
// }
