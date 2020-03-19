package main

import (
　　"crypto/rand"
　　"crypto/rsa"
　　"crypto/x509"
　　"crypto/x509/pkix"
　　"encoding/pem"
　　"math/big"
　　"net"
　　"os"
　　"time"
)
//生成SSL证书和秘钥
func main() {
　　max := new(big.Int).Lsh(big.NewInt(1), 128)
　　serialNumber, _ := rand.Int(rand.Reader, max)
//证书的标题
　　subject := pkix.Name{
　　　　Organization:　　　 []string{"Manning Publications Co."},
　　　　OrganizationalUnit: []string{"Books"},
　　　　CommonName:　　　　 "Go Web Programming",
　　}
//certificate结构对证书配置
　　template := x509.Certificate{
//证书序号（CA分发唯一
	　　　SerialNumber: serialNumber,
//证书的标题
　　　　Subject:　　　subject,

　　　　NotBefore:　　time.Now(),
　　　　NotAfter:　　 time.Now().Add(365 * 24 * time.Hour),
//表示用于服务器身份验证操作
　　　　KeyUsage:　　 x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
　　　　ExtKeyUsage:　[]x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
//在哪个IP上运行
		IPAddresses:　[]net.IP{net.ParseIP("127.0.0.1")},
　　}
//生成RSA私钥(结构包含公钥)
　　pk, _ := rsa.GenerateKey(rand.Reader, 2048)
//创建SSL证书
　　derBytes, _ := x509.CreateCertificate(rand.Reader, &template,
　　➥&template, &pk.PublicKey, pk)
//编码至cert.pem中
	certOut, _ := os.Create("cert.pem")
　　pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
　　certOut.Close()
//再编码保存到key.pem
　　keyOut, _ := os.Create("key.pem")
　　pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes:
　　➥x509.MarshalPKCS1PrivateKey(pk)})
　　keyOut.Close()
}