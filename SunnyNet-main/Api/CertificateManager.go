package Api

import "C"
import (
	"github.com/qtgolang/SunnyNet/public"
	"github.com/qtgolang/SunnyNet/src/Certificate"
)

// CreateCertificate 创建 证书管理器 对象
func CreateCertificate() int {
	return Certificate.CreateCertificate()
}

// RemoveCertificate 释放 证书管理器 对象
func RemoveCertificate(Context int) {
	Certificate.RemoveCertificate(Context)
}

// LoadP12Certificate 证书管理器 载入p12证书
func LoadP12Certificate(Context int, Name, Password string) bool {
	Certificate.Lock.Lock()
	defer Certificate.Lock.Unlock()
	c := Certificate.LoadCertificateContext(Context)
	if c == nil {
		return false
	}
	return c.LoadP12Certificate(Name, Password)
}

// LoadX509KeyPair 证书管理器 载入X509证书2
func LoadX509KeyPair(Context int, CaPath, KeyPath string) bool {
	Certificate.Lock.Lock()
	defer Certificate.Lock.Unlock()
	c := Certificate.LoadCertificateContext(Context)
	if c == nil {
		return false
	}
	return c.LoadX509KeyPair(CaPath, KeyPath)
}

// LoadX509Certificate 证书管理器 载入X509证书1
func LoadX509Certificate(Context int, Host, CA, KEY string) bool {
	Certificate.Lock.Lock()
	defer Certificate.Lock.Unlock()
	c := Certificate.LoadCertificateContext(Context)
	if c == nil {
		return false
	}
	return c.LoadX509Certificate(Host, CA, c.PrivateKey)
}

// SetInsecureSkipVerify 证书管理器 设置跳过主机验证
func SetInsecureSkipVerify(Context int, b bool) bool {
	Certificate.Lock.Lock()
	defer Certificate.Lock.Unlock()
	c := Certificate.LoadCertificateContext(Context)
	if c == nil {
		return false
	}
	return c.SetInsecureSkipVerify(b)
}

// SetServerName 证书管理器 设置ServerName
func SetServerName(Context int, name string) bool {
	Certificate.Lock.Lock()
	defer Certificate.Lock.Unlock()
	c := Certificate.LoadCertificateContext(Context)
	if c == nil {
		return false
	}
	return c.SetServerName(name)
}

// GetServerName 证书管理器 取ServerName
func GetServerName(Context int) uintptr {
	Certificate.Lock.Lock()
	defer Certificate.Lock.Unlock()
	c := Certificate.LoadCertificateContext(Context)
	if c == nil {
		return 0
	}
	return public.PointerPtr(c.GetServerName())
}

// AddCertPoolPath 证书管理器 设置信任的证书 从 文件
func AddCertPoolPath(Context int, cer string) bool {
	Certificate.Lock.Lock()
	defer Certificate.Lock.Unlock()
	c := Certificate.LoadCertificateContext(Context)
	if c == nil {
		return false
	}
	return c.AddCertPoolPath(cer)
}

// AddCertPoolText 证书管理器 设置信任的证书 从 文本
func AddCertPoolText(Context int, cer string) bool {
	Certificate.Lock.Lock()
	defer Certificate.Lock.Unlock()
	c := Certificate.LoadCertificateContext(Context)
	if c == nil {
		return false
	}
	return c.AddCertPoolText(cer)
}

// AddClientAuth 证书管理器 设置ClientAuth
func AddClientAuth(Context, val int) bool {
	Certificate.Lock.Lock()
	defer Certificate.Lock.Unlock()
	c := Certificate.LoadCertificateContext(Context)
	if c == nil {
		return false
	}
	return c.AddClientAuth(val)
}

// CreateCA 证书管理器 创建证书
func CreateCA(Context int, Country, Organization, OrganizationalUnit, Province, CommonName, Locality string, bits, NotAfter int) bool {
	Certificate.Lock.Lock()
	defer Certificate.Lock.Unlock()
	c := Certificate.LoadCertificateContext(Context)
	if c == nil {
		return false
	}
	return c.CreateCA(
		Country,
		Organization,
		OrganizationalUnit,
		Province,
		CommonName,
		Locality,
		bits,
		NotAfter)
}

// ExportCA 证书管理器 导出证书
func ExportCA(Context int) uintptr {
	Certificate.Lock.Lock()
	defer Certificate.Lock.Unlock()
	c := Certificate.LoadCertificateContext(Context)
	if c == nil {
		return 0
	}
	return public.PointerPtr(c.ExportCA())
}

// ExportKEY 证书管理器 导出私钥
func ExportKEY(Context int) uintptr {
	Certificate.Lock.Lock()
	defer Certificate.Lock.Unlock()
	c := Certificate.LoadCertificateContext(Context)
	if c == nil {
		return 0
	}
	return public.PointerPtr(c.ExportKEY())
}

// ExportPub 证书管理器 导出公钥
func ExportPub(Context int) uintptr {
	Certificate.Lock.Lock()
	defer Certificate.Lock.Unlock()
	c := Certificate.LoadCertificateContext(Context)
	if c == nil {
		return 0
	}
	return public.PointerPtr(c.ExportPub())
}

// GetCommonName 证书管理器 导出为P12
func GetCommonName(Context int) uintptr {
	Certificate.Lock.Lock()
	defer Certificate.Lock.Unlock()
	c := Certificate.LoadCertificateContext(Context)
	if c == nil {
		return 0
	}
	return public.PointerPtr(c.GetCommonName())
}

// ExportP12 证书管理器 导出为P12
func ExportP12(Context int, path, pass string) bool {
	Certificate.Lock.Lock()
	defer Certificate.Lock.Unlock()
	c := Certificate.LoadCertificateContext(Context)
	if c == nil {
		return false
	}
	return c.ExportP12(path, pass)
}
