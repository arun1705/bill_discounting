'use strict';

var cors = require('cors');
var path = require('path');
var request = require('request');
var cors = require('cors');



const generateInvoice = require('./functions/invoice');

const readInvoiceDetails = require('./functions/readInvoice.js');


module.exports = router => {


    router.post('/generateInvoice', cors(), (req, res) => {
        console.log("entering generateInvoice function in functions");
        const invoiceID = req.body.invoiceID;
        console.log(invoiceID);
        const goods = req.body.goods;
        const quantity = req.body.quantity;
        console.log(quantity);
        const duedate = req.body.dueDate;
        console.log(duedate);
        const invoicedate = req.body.invoiceDate;
        console.log(invoicedate);
        const to = req.body.to;
        console.log(to);
        const unitprice = req.body.unitPrice
        console.log(unitprice);
        const totalprice = req.body.totalPrice
        console.log(totalprice);


        if (!invoiceID.trim() || !goods.trim() || !quantity.trim() || !duedate.trim() || !invoicedate.trim() || !to.trim() || !unitprice.trim() || !totalprice.trim() || !invoiceID || !goods || !quantity || !duedate || !invoicedate || !to || !unitprice || !totalprice) {

            res.status(400).json({
                message: 'Invalid Request !'
            });

        } else {
            var value = {
                invoiceID: invoiceID,
                goods: goods,
                quantity: quantity,
                duedate: duedate,
                invoicedate: invoicedate,
                to: to,
                unitprice: unitprice,
                totalprice: totalprice
            }
            generateInvoice.generateInvoice(value)

                .then(result => {
                    16
                    res.status(result.status).json({
                        message: result.message,

                    })
                })

                .catch(err => res.status(err.status).json({
                    message: err.message
                }));
        }
    });


    router.post('/getInvoiceDetails', cors(), (req, res) => {

        const to = req.body.to;
        console.log(to);
        if (!to) {
            res.status(400).json({
                message: 'Invalid Request !'
            });
        } else {

            readInvoiceDetails.readInvoiceDetails(to)
                .then(function (result) {
                    return res.status(200).json({
                        "message": result.message
                    });
                })
                .catch(err => res.status(err.status).json({
                    message: err.message
                }));
        }

    });


    router.post('/financeInvoice', cors(), (req, res) => {
        console.log("entering financeInvoice function in functions");
        const invoiceID = req.body.invoiceID;
        console.log(invoiceID);
        const goods = req.body.goods;
        const quantity = req.body.quantity;
        console.log(quantity);
        const duedate = req.body.dueDate;
        console.log(duedate);
        const invoicedate = req.body.invoiceDate;
        console.log(invoicedate);
        const to = req.body.to;
        console.log(to);
        const unitprice = req.body.unitPrice
        console.log(unitprice);
        const discount = req.body.discount;
        console.log(discount);
        const totalprice = req.body.totalPrice;
        console.log(totalprice);
        const discountPrice = req.body.discountPrice;
        console.log(discountPrice);
        const status = req.body.status;
        console.log(status);


        if (!invoiceID || !goods || !quantity || !duedate || !invoicedate || !unitprice || !discount || !discountPrice || !totalprice || !to || !status || !invoiceID.trim() || !goods.trim() || !quantity.trim() || !duedate.trim() || !invoicedate.trim() || !unitprice.trim() || !discount.trim() || !discountPrice.trim() || !totalprice.trim() || !to.trim() || !status.trim()) {

            res.status(400).json({
                message: 'Invalid Request !'
            });

        } else {
            var value = {
                invoiceID: invoiceID,
                goods: goods,
                quantity: quantity,
                duedate: duedate,
                invoicedate: invoicedate,
                unitprice: unitprice,
                discount: discount,
                totalprice: totalprice,
                discountPrice: discountPrice,
                to: to,
                status: status
            }
            generateInvoice.financeInvoice(value)

                .then(result => {

                    res.status(result.status).json({
                        message: result.message,

                    })
                })

                .catch(err => res.status(err.status).json({
                    message: err.message
                }));
        }
    });


    router.post('/getAllDetails', cors(), (req, res) => {

        const startInvoiceId = req.body.startInvoiceId;
        const endInvoiceId = req.body.endInvoiceId

        if (!startInvoiceId || !endInvoiceId) {
            res.status(400).json({
                message: 'Invalid Request !'
            });
        } else {

            readInvoiceDetails.readAllInvoiceDetails(startInvoiceId, endInvoiceId)
                .then(function (result) {
                    return res.json({
                        "message": result
                    });
                })
                .catch(err => res.status(err.status).json({
                    message: err.message
                }));
        }

    });

    router.post("/Login", cors(), (req, res) => {
        var username = req.body.username;
        var password = req.body.password;
        console.log(req.body)
        console.log(username);
        if (username === "seller@discount.com") {
            res.send({
                "message": "Login Successful",
                "status": true,
                "userType": "seller"
            })
        } else if (username == "buyer@discount.com") {
            res.send({
                "message": "Login Successful",
                "status": true,
                "userType": "buyer"
            })
        } else if (username == "bank@discount.com") {
            res.send({
                "message": "Login Successful",
                "status": true,
                "userType": "banker"
            })
        }
    })
}
