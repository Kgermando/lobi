package helpers

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/lobi/database"
	"github.com/kgermando/lobi/models"
)

func SendVerificationMail(c *fiber.Ctx, email string, pass string) (string, error) {
	host := os.Getenv("SMTP_HOST")
	port := os.Getenv("SMTP_PORT")
	from := os.Getenv("SMTP_MAIL")
	password := os.Getenv("SMTP_PASSWORD")

	subject := "Email de verification"
	auth := smtp.PlainAuth("", from, password, host)

	// verificationLinkBase := "http://localhost:3000/web/users/verify-compte?email="
	verificationLinkBase := "https://gouvdev-rdc.net/web/users/verify-compte?email="

	// verificationKey := fmt.Sprintf(
	// 	"%x", sha256.Sum256([]byte(email + "-" + uuid.New().String())[:]),
	// )

	body := fmt.Sprintf(`
	<html>
		<head>
			<meta charset="UTF-8">
			<title>Validation d'adresse mail</title>
			<style>
				body {
					margin: 0;
					padding: 20px;
				}
				
				.container {
					width: 600px;
					margin: 0 auto;
				}
				
				.footer {
					text-align: center;
					font-size: 12px;
				}
				h1 {
					font-family: Arial, sans-serif;
					font-size: 18px;
				}
				
				p {
					font-family: Arial, sans-serif;
					font-size: 14px;
					line-height: 1.5;
				}
				h1 {
					color: #333;
				}
				
				p {
					color: #666;
				}
				
				a {
					display: inline-block;
					padding: 10px 20px;
					border: 2px solid #000;
					border-radius: 5px;
					color: #000;
					text-decoration: none;
					transition: all 0.2s ease-in-out;
				}
				
				a:hover {
					background-color: transparent;
					color: #000;
					border-color: #f00;
				}
			</style>
		</head>
		<body>
			<div class="container">
				<h1>Validation de votre compte GouvDev</h1>
				<p>Bonjour </p>
				<p>Cliquez sur le lien ci-après pour valider votre adresse mail.</p>
				<a href="%v%v" target="_blank">Valider</a>
				<p>Ce lien expire dans 24 heures.</p>
				<p>Cordialement,</p>
				<p>L'équipe GouvDev</p>
			</div>
			<div class="footer">
				<p>Copyright &copy; 2024 GouvDev</p>
			</div>
		</body> 
	</html>
	`, verificationLinkBase, email)

	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	cc, err := smtp.Dial(host + ":" + port)
	if err != nil {
		return "", err
	}

	cc.StartTLS(tlsconfig)

	if err = cc.Auth(auth); err != nil {
		fmt.Println(err)
		return "", err
	}

	if err = cc.Mail(from); err != nil {
		return "", err
	}

	if err = cc.Rcpt(email); err != nil {
		return "", err
	}

	w, err := cc.Data()
	if err != nil {
		return "", err
	}

	_, err = w.Write([]byte(
		fmt.Sprintf("MIME-Version: %v\r\n", "1.0") +
			fmt.Sprintf("Content-type: %v\r\n", "text/html; charset=UTF-8") +
			fmt.Sprintf("From: %v\r\n", from) +
			fmt.Sprintf("To: %v\r\n", email) +
			fmt.Sprintf("Subject: %v\r\n", subject) +
			fmt.Sprintf("%v\r\n", body),
	))

	if err != nil {
		return "", err
	}

	db := database.DB
	user := new(models.User)

	user.Email = email
	user.Password = pass
	user.EmailVerified = false
	user.IsActive = false
	user.Role = "User"

	if err := db.Create(&user).Error; err != nil {
		fmt.Println("Une erreur s'est produite", err.Error())
		c.Redirect("/web/auth/register")
	}

	err = w.Close()
	if err != nil {
		return "", err
	}

	cc.Quit()

	return fmt.Sprintf("%v%v", verificationLinkBase, email), nil

}
