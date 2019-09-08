/**
 * Responds to any HTTP request.
 *
 * @param {!express:Request} req HTTP request context.
 * @param {!express:Response} res HTTP response context.
 */

 exports.getUniform = (req, res) => {
  res.json({
    "data": {
      "metadata": {
        "app_name": "IDSchool",
        "app_id": 111,
        "title": "Đồng phục học sinh năm 2019",
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
            "content": "Bộ mùa đông: 600.000 vnđ",
            "status": "paid",
          },
          {
            "type": "text",
            "style": "paragraph",
            "content": "Bộ mùa hè: 500.000 vnđ",
            "status": "paid",
          },
          {
            "label": "Tùy chọn",
            "type": "checkbox",
            "display_type": "inline",
            "name": "total_payment",
            "options": [
              {
                "label": "Mùa đông",
                "value": 600000
              },
              {
                "label": "Mùa hè",
                "value": 500000
              }
            ]
          }
        ]
      }
    }
  });
};