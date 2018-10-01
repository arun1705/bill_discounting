
'use strict';

var bcSdk = require('../fabcar/invoke.js');


/**
 * A module that will generate Invoice and store into the blockchain!
 * @module generateInvoice
 */
/** Add Data into blockchain.*/
exports.generateInvoice = (value) => {
    return new Promise((resolve, reject) => {

        value.status = "invoiceRaised"
        bcSdk.generateInvoice(value)
            .then(function (result) {
            
                resolve({
                    "status": 200,
                    "message": result
                });

            })
            .catch(err => {

                reject({
                    "status": 500,
                    "message": 'Something went wrong please try again later!!'
                });

            });
    });
}

/**
 * A module that will generate Invoice and store into the blockchain!
 * @module financeInvoice
 */
/** Add Data into blockchain.*/
exports.financeInvoice = (value) => {
    return new Promise((resolve, reject) => {

        if (value.status=="invoiceRaised"){
            value.status="approvalPending"
        }else if(value.status=="approvalPending"){
            value.status="approved"
        }else if(value.status="approved"){
            value.status="financeReleased"
        }else if(value.status="rejected"){
            value.status="rejected"
        }
        bcSdk.financeInvoice(value)
            .then(function (result) {
            
                resolve({
                    "status": 200,
                    "message": result
                });

            })
            .catch(err => {

                reject({
                    "status": 500,
                    "message": 'Something went wrong please try again later!!'
                });

            });
    });
}