package email

import (
	"context"
	"fmt"
	"github.com/ITu-CloudWeGo/itu_rpc_auth/rpc/config"
	"github.com/redis/go-redis/v9"
	"gopkg.in/gomail.v2"
	"log"
	"math/rand"
	"time"
)

var (
	globalConfig *config.Config
	redisClient  *redis.Client
)

func init() {
	var err error
	globalConfig, err = config.GetConfig()
	if err != nil {
		panic(fmt.Sprintf("failed to load config: %v", err))
	}

	redisClient = redis.NewClient(&redis.Options{
		Addr:     globalConfig.Redis.Address,
		Password: globalConfig.Redis.Password,
		DB:       globalConfig.Redis.DB,
	})
}

// GenerateCaptcha 生成4位随机验证码
func GenerateCaptcha() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(9000) + 1000
}

// SendCaptchaEmail 发送验证码邮件
func SendCaptchaEmail(to string) (int, error) {
	captcha := GenerateCaptcha()
	subject := "欢迎使用ITU"
	from := "1356918183@qq.com"
	body := fmt.Sprintf("您的验证码是：%s。密码五分钟过期，请尽快完成验证。", captcha)

	// 发送邮件
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	d := gomail.NewDialer("smtp.qq.com", 587, "1356918183@qq.com", "ckujnbafgznlibfe") // SMTP服务器配置

	if err := d.DialAndSend(m); err != nil {
		log.Printf("发送邮件失败: %v", err)
		return 0, err
	}

	ctx := context.Background()
	err := redisClient.Set(ctx, to, captcha, 5*time.Minute).Err()
	if err != nil {
		log.Printf("存储验证码到Redis失败: %v", err)
		return 0, err
	}

	return captcha, nil
}
