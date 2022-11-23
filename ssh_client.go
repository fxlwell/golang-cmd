package cmd

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

type SSHConf struct {
	User     string
	Password string
	Host     string
	Port     int
	CertFile string
	Timeout  time.Duration
}

type SSHClient struct {
	client     *ssh.Client
	sftpClient *sftp.Client
	config     *ssh.ClientConfig
	addr       string
}

func SSHConfigWithPassword(user, password string, timeout time.Duration) *ssh.ClientConfig {
	return &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         timeout,
	}
}

func SSHConfigWithCertFile(user, certFile string, timeout time.Duration) (*ssh.ClientConfig, error) {
	if content, err := ioutil.ReadFile(certFile); err != nil {
		return nil, err
	} else {
		if signer, err := ssh.ParsePrivateKey(content); err != nil {
			return nil, err
		} else {
			return &ssh.ClientConfig{
				User: user,
				Auth: []ssh.AuthMethod{
					ssh.PublicKeys(signer),
				},
				HostKeyCallback: ssh.InsecureIgnoreHostKey(),
				Timeout:         timeout,
			}, nil
		}
	}
	return nil, fmt.Errorf("")
}

func NewSSHClient(conf *SSHConf) (*SSHClient, error) {
	var (
		config  *ssh.ClientConfig
		client  *ssh.Client
		timeout time.Duration
		err     error
	)

	if conf.Timeout > 0 {
		timeout = conf.Timeout
	} else {
		timeout = time.Second * 30 //default
	}

	if len(conf.Password) == 0 && len(conf.CertFile) != 0 {
		if config, err = SSHConfigWithCertFile(conf.User, conf.CertFile, timeout); err != nil {
			return nil, err
		}
	} else {
		config = SSHConfigWithPassword(conf.User, conf.Password, timeout)
	}

	address := fmt.Sprintf("%s:%d", conf.Host, conf.Port)

	client, err = ssh.Dial("tcp", address, config)
	if err != nil {
		return nil, err
	}

	sftpClient, err = sftp.NewClient(sshClient)
	if err != nil {
		return nil, err
	}

	return &SSHClient{
		client: client,
		config: config,
		addr:   address,
	}, nil
}

func (sshc *SSHClient) RunCommand(cmd string) error {
	sshc.client.NewSection()
	defer s.Close()
}

func (sshc *SSHClient) SendFile(localPath, remotePath string) error {
	ssh.Once.Do(func() {
		sftpClient, err = sftp.NewClient(sshClient)
		if err != nil {
			return nil, err
		}

	})
}

func (sshc *SSHClient) ReadFile(remotePath, localPath string) error {

}

func (sshc *SSHClient) SendDir(localDir, remoteDir string) error {

}

func (sshc *SSHClient) ReadDir(remoteDir, localDir string) error {

}

func (sshc *SSHClient) Close() error {

}
