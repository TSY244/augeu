package api

import (
	middleware2 "augeu/backEnd/internal/pkg/web/middleware"
	"augeu/backEnd/internal/utils/consts/web"
	"augeu/backEnd/internal/utils/utils"
	"augeu/public/pkg/swaggerCore/models"
	"augeu/public/pkg/swaggerCore/restapi/operations"
	"augeu/public/util/convert"
	"encoding/json"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"net/http"
)

func (apiManager *ApiManager) GetFileReportPostApiHandlerFunc() operations.PostGetFileReportHandlerFunc {
	return func(params operations.PostGetFileReportParams) middleware.Responder {
		apiName := "GetFileReportPostApi"
		if resp := middleware2.CheckAgentRole(params.HTTPRequest, apiManager.s); resp != nil {
			return resp
		}
		data := params.Data
		if data == nil {
			return operations.NewPostGetFileReportBadRequest().WithPayload(&models.BadRequestError{
				Code:    convert.Int64P(int64(operations.PostGetFileReportBadRequestCode)),
				Message: convert.StrPtr("param is nil"),
			})
		}
		target := data.Target
		if target == nil {
			return operations.NewPostGetFileReportBadRequest().WithPayload(&models.BadRequestError{
				Code:    convert.Int64P(int64(operations.PostGetFileReportBadRequestCode)),
				Message: convert.StrPtr("param is nil"),
			})
		}
		apiKey := apiManager.s.Config.CoreConfig.WeiBuConf.ApiKey
		if apiKey == "" {
			return operations.NewPostGetFileReportInternalServerError().WithPayload(&models.ActionFailure{
				From:    utils.StrP(apiName),
				Reason:  utils.StrP(web.InternalError),
				Success: web.Fail,
			})
		}
		url := fmt.Sprintf("https://api.threatbook.cn/v3/file/report?apikey=%s&sandbox_type=win10_1903_enx64_office2016&resource=%s&query_fields=summary&query_fields=multiengines", apiKey, *target)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return operations.NewPostGetFileReportInternalServerError().WithPayload(&models.ActionFailure{
				From:    utils.StrP(apiName),
				Reason:  utils.StrP(web.InternalError),
				Success: web.Fail,
			})
		}
		weiBuResp, err := http.DefaultClient.Do(req)
		if err != nil {
			return operations.NewPostGetFileReportInternalServerError().WithPayload(&models.ActionFailure{
				From:    utils.StrP(apiName),
				Reason:  utils.StrP(web.InternalError),
				Success: web.Fail,
			})
		}
		// json -> go struct
		defer weiBuResp.Body.Close()
		var weiBuRespStruct weiBuResponse
		var databytes []byte
		size, err := weiBuResp.Body.Read(databytes)
		if err != nil || size == 0 {
			return operations.NewPostGetFileReportInternalServerError().WithPayload(&models.ActionFailure{
				From:    utils.StrP(apiName),
				Reason:  utils.StrP(web.InternalError),
				Success: web.Fail,
			})
		}

		err = json.Unmarshal(databytes, &weiBuRespStruct)
		if err != nil {
			return operations.NewPostGetFileReportInternalServerError().WithPayload(&models.ActionFailure{
				From:    utils.StrP(apiName),
				Reason:  utils.StrP(web.InternalError),
				Success: web.Fail,
			})
		}
		if weiBuRespStruct.ResponseCode != 200 {
			return operations.NewPostGetFileReportInternalServerError().WithPayload(&models.ActionFailure{
				From:    utils.StrP(apiName),
				Reason:  utils.StrP(web.InternalError),
				Success: web.Fail,
			})
		}
		return operations.NewPostGetFileReportOK().WithPayload(&models.GetFileReportResponse{
			Data:         convert2WeiBuData(&weiBuRespStruct),
			ResponseCode: convert.Int64P(int64(weiBuRespStruct.ResponseCode)),
		})
	}
}

// weiBuResponse 根结构体
type weiBuResponse struct {
	ResponseCode int       `json:"response_code"`
	Data         weiBuData `json:"data"`
	VerboseMsg   string    `json:"verbose_msg"`
}

// weiBuData 包含各种信息的结构体
type weiBuData struct {
	Summary      weiBuSummary      `json:"summary"`
	Multiengines weiBuMultiengines `json:"multiengines"`
	Static       Static            `json:"static"`
	Signature    []weiBuSignature  `json:"signature"`
	Dropped      []weiBuDropped    `json:"dropped"`
	Pstree       weiBuPstree       `json:"pstree"`
	Network      weiBuNetwork      `json:"network"`
	Strings      weiBuStrings      `json:"strings"`
	Permalink    string            `json:"permalink"`
}

// weiBuSummary 概要信息结构体
type weiBuSummary struct {
	ThreatLevel     string   `json:"threat_level"`
	MalwareType     string   `json:"malware_type"`
	MalwareFamily   string   `json:"malware_family"`
	IsWhitelist     bool     `json:"is_whitelist"`
	SubmitTime      string   `json:"submit_time"`
	FileName        string   `json:"file_name"`
	FileType        string   `json:"file_type"`
	SampleSha256    string   `json:"sample_sha256"`
	MD5             string   `json:"md5"`
	SHA1            string   `json:"sha1"`
	Tag             Tag      `json:"tag"`
	ThreatScore     int      `json:"threat_score"`
	SandboxType     string   `json:"sandbox_type"`
	SandboxTypeList []string `json:"sandbox_type_list"`
	MultiEngines    string   `json:"multi_engines"`
}

// Tag 标签结构体
type Tag struct {
	S []string `json:"s"`
	X []string `json:"x"`
}

// weiBuMultiengines 反病毒扫描引擎检测结果结构体
type weiBuMultiengines struct {
	Result   Result `json:"result"`
	ScanTime string `json:"scan_time"`
}

// Result 反病毒扫描引擎具体结果结构体
type Result struct {
	Kaspersky string `json:"Kaspersky"`
	Microsoft string `json:"Microsoft"`
}

// Static 静态信息结构体
type Static struct {
	Details weiBuStaticDetails `json:"details"`
	Basic   Basic              `json:"basic"`
}

// weiBuStaticDetails PE 文件静态信息结构体
type weiBuStaticDetails struct {
	PeVersionInfo []interface{}          `json:"pe_version_info"`
	PeSections    []interface{}          `json:"pe_sections"`
	PeSignatures  map[string]interface{} `json:"pe_signatures"`
	PeImports     []interface{}          `json:"pe_imports"`
	PeResources   []interface{}          `json:"pe_resources"`
	Tag           []interface{}          `json:"tag"`
	PeDetect      map[string]interface{} `json:"pe_detect"`
	PeBasic       map[string]interface{} `json:"pe_basic"`
	PeExports     []interface{}          `json:"pe_exports"`
}

// Basic 文件基本信息结构体
type Basic struct {
	SHA1     string `json:"sha1"`
	SHA256   string `json:"sha256"`
	FileType string `json:"file_type"`
	FileName string `json:"file_name"`
	SSDeep   string `json:"ssdeep"`
	FileSize int    `json:"file_size"`
	MD5      string `json:"md5"`
}

// weiBuSignature 行为签名结构体
type weiBuSignature struct {
	Severity    int                    `json:"severity"`
	References  []interface{}          `json:"references"`
	SigClass    string                 `json:"sig_class"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Markcount   int                    `json:"markcount"`
	Marks       []interface{}          `json:"marks"`
	Families    []interface{}          `json:"families"`
	AttckID     string                 `json:"attck_id"`
	AttckInfo   map[string]interface{} `json:"attck_info"`
}

// weiBuDropped 释放行为结构体
type weiBuDropped struct {
	SHA1     string        `json:"sha1"`
	URLs     []string      `json:"urls"`
	SHA256   string        `json:"sha256"`
	Size     int           `json:"size"`
	Filepath string        `json:"filepath"`
	Name     string        `json:"name"`
	CRC32    string        `json:"crc32"`
	SSDeep   string        `json:"ssdeep"`
	Type     string        `json:"type"`
	Yara     []interface{} `json:"yara"`
	MD5      string        `json:"md5"`
}

// weiBuPstree 进程行为结构体
type weiBuPstree struct {
	Children    []Children       `json:"children"`
	ProcessName weiBuProcessName `json:"process_name"`
}

// Children 进程子结构体
type Children struct {
	Track       bool       `json:"track"`
	PID         int        `json:"pid"`
	ProcessName string     `json:"process_name"`
	CommandLine string     `json:"command_line"`
	FirstSeen   string     `json:"first_seen"`
	PPID        int        `json:"ppid"`
	Children    []Children `json:"children"`
}

// weiBuProcessName 进程名称信息结构体
type weiBuProcessName struct {
	En string `json:"en"`
	Cn string `json:"cn"`
}

// weiBuNetwork 网络行为结构体
type weiBuNetwork struct {
	Fingerprint []interface{} `json:"fingerprint"`
	TLS         []interface{} `json:"tls"`
	UDP         []interface{} `json:"udp"`
	DnsServers  []interface{} `json:"dns_servers"`
	HTTP        []interface{} `json:"http"`
	IRC         []interface{} `json:"irc"`
	SMTP        []interface{} `json:"smtp"`
	TCP         []interface{} `json:"tcp"`
	SMTPEx      []interface{} `json:"smtp_ex"`
	MITM        []interface{} `json:"mitm"`
	Hosts       []interface{} `json:"hosts"`
	DNS         []interface{} `json:"dns"`
	HTTPEx      []interface{} `json:"http_ex"`
	Domains     []interface{} `json:"domains"`
	DeadHosts   []interface{} `json:"dead_hosts"`
	ICMP        []interface{} `json:"icmp"`
	HTTPSEx     []interface{} `json:"https_ex"`
}

// weiBuStrings 字符串信息结构体
type weiBuStrings struct {
	SHA256 []string `json:"sha256"`
	Pcap   []string `json:"pcap"`
}

func convert2WeiBuData(src *weiBuResponse) *models.WeiBuData {
	// 转换 Signature
	signatures := make([]*models.WeiBuSignatureBase, 0, len(src.Data.Signature))
	for _, sig := range src.Data.Signature {
		severityStr := fmt.Sprintf("%d", sig.Severity)
		sigCopy := sig // 避免引用问题
		signatures = append(signatures, &models.WeiBuSignatureBase{
			Description: &sigCopy.Description,
			Severity:    &severityStr,
			SigClass:    &sigCopy.SigClass,
		})
	}

	toStringSlice := func(input []interface{}) []string {
		result := make([]string, 0, len(input))
		for _, item := range input {
			if str, ok := item.(string); ok {
				result = append(result, str)
			}
		}
		return result
	}

	// 构建 Network 结构
	network := &models.WeiBuNetwork{
		DeadHosts:   toStringSlice(src.Data.Network.DeadHosts),
		DNS:         toStringSlice(src.Data.Network.DNS),
		DNSServers:  toStringSlice(src.Data.Network.DnsServers),
		Domains:     toStringSlice(src.Data.Network.Domains),
		Fingerprint: toStringSlice(src.Data.Network.Fingerprint),
		Hosts:       toStringSlice(src.Data.Network.Hosts),
		HTTP:        toStringSlice(src.Data.Network.HTTP),
		HTTPEx:      toStringSlice(src.Data.Network.HTTPEx),
		HTTPSEx:     toStringSlice(src.Data.Network.HTTPSEx),
		Icmp:        toStringSlice(src.Data.Network.ICMP),
		Irc:         toStringSlice(src.Data.Network.IRC),
		Mitm:        toStringSlice(src.Data.Network.MITM),
		SMTP:        toStringSlice(src.Data.Network.SMTP),
		SMTPEx:      toStringSlice(src.Data.Network.SMTPEx),
		TCP:         toStringSlice(src.Data.Network.TCP),
		TLS:         toStringSlice(src.Data.Network.TLS),
		UDP:         toStringSlice(src.Data.Network.UDP),
	}

	// 转换 ThreatScore 为 *int64
	threatScore := int64(src.Data.Summary.ThreatScore)

	// 构建最终结构体
	return &models.WeiBuData{
		MalwareFamily: &src.Data.Summary.MalwareFamily,
		MultiEngines:  &src.Data.Summary.MultiEngines,
		Network:       network,
		Permalink:     &src.Data.Permalink,
		Signature:     &models.WeiBuSignature{Severity: signatures},
		Strings: &models.Strings{
			Pcap:   src.Data.Strings.Pcap,
			Sha256: src.Data.Strings.SHA256,
		},
		SubmitTime:  &src.Data.Summary.SubmitTime,
		Tag:         &models.WeiBuTag{S: src.Data.Summary.Tag.S, X: src.Data.Summary.Tag.X},
		ThreatLevel: &src.Data.Summary.ThreatLevel,
		ThreatScore: &threatScore,
		ThreatType:  &src.Data.Summary.MalwareType,
	}
}
