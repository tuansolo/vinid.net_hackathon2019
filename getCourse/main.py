from flask import jsonify
def get_course(request):
    """Responds to any HTTP request.
    Args:
        request (flask.Request): HTTP request object.
    Returns:
        The response text or any set of values that can be turned into a
        Response object using
        `make_response <http://flask.pocoo.org/docs/1.0/api/#flask.Flask.make_response>`.
    """
    request_json = request.get_json()
    if request.args and 'message' in request.args:
        return request.args.get('message')
    elif request_json and 'message' in request_json:
        return request_json['message']
    else:
        result = {
            "data": {
                "metadata": {
                    "app_name": "Đăng ký khóa học",
                    "app_id": 123456,
                    "title": "Đăng ký khóa học",
                    "submit_button": {
                        "label": "Đăng ký",
                        "background_color": "#6666ff",
                        "cta": "request",
                        "url": "https://us-central1-gentle-epoch-234103.cloudfunctions.net/getPaymentOptions"
                    },
                    "elements": [
                        {
                            "type": "text",
                            "style": "paragraph",
                            "content": "Hội họa - Tranh hiện đại (đã đăng ký)",
                            "status": "paid",
                        },
                        {
                            "label": "Các bộ môn năng khiếu",
                            "type": "checkbox",
                            "display_type": "inline",
                            "required": False,
                            "error": "Thông tin này là bắt buộc",
                            "name": "total_payment",
                            "placeholder": "Các bộ môn năng khiếu",
                            "options": [
                                {
                                    "label": "Bơi sải - 1.000.000",
                                    "value": 1000000
                                },
                                {
                                    "label": "Bơi ếch - 1.500.000",
                                    "value": 1500000
                                },
                                {
                                    "label": "Bơi bướm - 800.000",
                                    "value": 800000
                                },
                                {
                                    "label": "Bơi tự do - 1.200.000",
                                    "value": 1200000
                                },
                                {
                                    "label": "Piano - 2.000.000",
                                    "value": 2000000
                                },
                                {
                                    "label": "Guitar - 1.800.000",
                                    "value": 1800000
                                },
                                {
                                    "label": "Violon - 1.600.000",
                                    "value": 1600000
                                }
                            ]
                        },
                        {
                            "type": "input",
                            "input_type": "text",
                            "label": "Ghi chú",
                            "required": False,
                            "name": "note",
                            "error": "Vui lòng nhập đủ thông tin",
                            "placeholder": "Hãy ghi lại mong muốn của bạn"
                        }
                    ]
                }
            }
        }
        return jsonify(result)
