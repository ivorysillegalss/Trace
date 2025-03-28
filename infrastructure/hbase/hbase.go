package hbase

import (
	"context"
	"log"

	"github.com/tsuna/gohbase"
	"github.com/tsuna/gohbase/filter"
	"github.com/tsuna/gohbase/hrpc"
)

type Client interface {
	// 增加记录
	Insert(ctx context.Context, table string, key string, values map[string]map[string][]byte) (strValue string, Exists bool, err error)
	// 获取整行记录
	GetRow(ctx context.Context, table string, row string) (strValue string, Exists bool, err error)
	// 获取特定行记录
	GetCell(ctx context.Context, table string, key string, values map[string][]string) (strValue string, Exists bool, err error)
	// TODO 跟应用自带包进行解耦
	// +筛选条件获取记录
	GetCellWithFilter(ctx context.Context, table string, key string, values map[string][]string, filter filter.Filter) (strValue string, Exists bool, err error)
	// +筛选条件进行扫描
	ScanWithFilter(ctx context.Context, table string, filter filter.Filter) map[string]int64
}

type hbaseClient struct {
	client gohbase.Client
}

// GetCell implements Client.
func (h *hbaseClient) GetCell(ctx context.Context, table string, key string, values map[string][]string) (strValue string, Exists bool, err error) {
	getRequest, err := hrpc.NewGetStr(ctx, table, key, hrpc.Families(values))
	if err != nil {
		log.Fatal("Hbase: get value error!\n")
	}
	getResp, err := h.client.Get(getRequest)
	return getResp.String(), *getResp.Exists, err
}

// GetCellWithFilter implements Client.
func (h *hbaseClient) GetCellWithFilter(ctx context.Context, table string, key string, values map[string][]string, filter filter.Filter) (strValue string, Exists bool, err error) {
	getRequest, err := hrpc.NewGetStr(ctx, table, key, hrpc.Families(values), hrpc.Filters(filter))
	if err != nil {
		log.Fatal("Hbase: get value error!\n")
	}
	getResp, err := h.client.Get(getRequest)
	return getResp.String(), *getResp.Exists, err
}

// ScanWithFilter implements Client.
func (h *hbaseClient) ScanWithFilter(ctx context.Context, table string, filter filter.Filter) map[string]int64 {
	scanRequest, err := hrpc.NewScanStr(ctx, table, hrpc.Filters(filter))
	if err != nil {
		log.Fatal("Hbase: get value error!\n")
	}
	scanResp := h.client.Scan(scanRequest)
	return scanResp.GetScanMetrics()

}

// Insert implements Client.
func (h *hbaseClient) Insert(ctx context.Context, table string, key string, values map[string]map[string][]byte) (strValue string, Exists bool, err error) {
	putRequest, err := hrpc.NewPutStr(ctx, table, key, values)
	if err != nil {
		log.Fatal("Hbase: insert error!\n")
	}
	resp, err := h.client.Put(putRequest)
	return resp.String(), *resp.Exists, err
}

func (h *hbaseClient) GetRow(ctx context.Context, table string, row string) (strValue string, Exists bool, err error) {
	getRequest, err := hrpc.NewGetStr(ctx, table, row)
	if err != nil {
		log.Fatal("Hbase: get value error!\n")
	}
	getResp, err := h.client.Get(getRequest)
	return getResp.String(), *getResp.Exists, err
}

func NewHbaseClient(zkquorum string, args ...any) Client {
	// client := gohbase.NewClient("localhost")
	client := gohbase.NewClient(zkquorum)
	return &hbaseClient{client: client}
}
