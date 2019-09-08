/**
 * Responds to any HTTP request.
 *
 * @param {!express:Request} req HTTP request context.
 * @param {!express:Response} res HTTP response context.
 */
exports.getPaymentOptions = (req, res) => {
  console.log('req.body: ', req.body);
  totalPayment = 0;

  if(req.body.total_payment) {
    req.body.total_payment.map(value => totalPayment += parseInt(value))
  }
  console.log('totalPayment: ', totalPayment);
  res.json({
    "data": {
      "metadata": {
        "app_name": "IDSchool",
        "app_id": 111,
        "title": "Thanh Toán",
        "submit_button": {
          "label": "Thanh toán",
          "background_color": "#6666ff",
          "cta": "request",
          "url": "https://us-central1-gentle-epoch-234103.cloudfunctions.net/gateway"
        },
        "elements": [
          {
            "type": "text",
            "style": "heading",
            "content": `Tổng tiền: ${totalPayment}`,
          },
          {
            "label": "Phương thức thanh toán",
            "type": "radio",
            "display_type": "inline",
            "name": "payment_method",
            "options": [
              {
                "label": "Ví điện tử",
                "value": "wallet"
              },
              {
                "label": "Thẻ ATM",
                "value": "atm"
              },
              {
                "label": "Visa",
                "value": "visa"
              }
            ]
          }
        ]
      }
    }
  });
};
