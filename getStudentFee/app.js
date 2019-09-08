/**
 * Responds to any HTTP request.
 *
 * @param {!express:Request} req HTTP request context.
 * @param {!express:Response} res HTTP response context.
 */
exports.getStudentFee = (req, res) => {
  res.json({
    "data": {
      "metadata": {
        "app_name": "IDSchool",
        "app_id": 111,
        "title": "Xem thông tin học phí",
        "submit_button": {
          "label": "Tiếp Tục",
          "background_color": "#6666ff",
          "cta": "request",
          "url": "https://us-central1-gentle-epoch-234103.cloudfunctions.net/getPaymentOptions"
        },
        "elements": [
          {
            "type": "text",
            "style": "paragraph",
            "content": "Năm 2: 20.000.000",
            "status": "paid",
          },
          {
            "type": "text",
            "style": "paragraph",
            "content": "Năm 1: 10.500.300",
            "status": "paid",
          },
          {
            "label": "Đóng học phí",
            "type": "checkbox",
            "display_type": "inline",
            "name": "total_payment",
            "options": [
              {
                "label": "Năm 3: 30.000.000",
                "value": 30000000
              },
              {
                "label": "Năm 4: 40.000.000",
                "value": 40000000
              }
            ]
          }
        ]
      }
    }
  });
};
