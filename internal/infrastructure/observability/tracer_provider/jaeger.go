package tracer_provider

import (
	"context"
	"log"

	"go.opentelemetry.io/contrib/propagators/aws/xray"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

/** Jaeger向けのOpenTelemetryTracerを初期化する */
func InitJaegerTracer(ctx context.Context, dsn string, resource *resource.Resource, insecure bool) func(context.Context) error {
	// gRPCプロトコルで 指定エンドポイントに対してトレースデータを送信するエクスポーターを作成
	var (
		exporter *otlptrace.Exporter
		err      error
	)
	if insecure {
		exporter, err = otlptracegrpc.New(
			ctx,
			otlptracegrpc.WithInsecure(),
			otlptracegrpc.WithEndpoint(dsn),
		)
	} else {
		exporter, err = otlptracegrpc.New(
			ctx,
			otlptracegrpc.WithEndpoint(dsn),
		)
	}
	if err != nil {
		log.Fatal(err)
	}

	// スパンが完成するまでキューイングするバッチスパンプロセッサーを作成
	bsp := sdktrace.NewBatchSpanProcessor(exporter,
		// 1000個までキューする
		sdktrace.WithMaxQueueSize(1000),
		// 1000個までまとめてエクスポートする
		sdktrace.WithMaxExportBatchSize(1000),
	)

	// トレーシングする際の設定を行う トレーサープロバイダー を作成
	tp := sdktrace.NewTracerProvider(
		// 常にspanに付与させるリソースを設定
		sdktrace.WithResource(resource),
		// スパンのID生成器を設定
		sdktrace.WithIDGenerator(xray.NewIDGenerator()),
	)

	// トレーサープロバイダーに バッチスパンプロセッサー を登録
	tp.RegisterSpanProcessor(bsp)

	// OpenTelemetry(グローバル) にトレーサープロバイダーを登録
	otel.SetTracerProvider(tp)

	// アプリ終了前にgracefulにエクスポーターをシャットダウンするための関数を返す
	return bsp.Shutdown
}
