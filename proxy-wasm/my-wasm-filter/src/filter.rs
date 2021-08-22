use log::info;
use proxy_wasm::traits::*;
use proxy_wasm::types::*;
use std::time::Duration;

pub struct HelloWorldFilterRoot;

impl Context for HelloWorldFilterRoot {}

impl RootContext for HelloWorldFilterRoot {
    fn get_type(&self) -> Option<ContextType> {
        Some(ContextType::HttpContext)
    }

    fn create_http_context(&self, _: u32) -> Option<Box<dyn HttpContext>> {
        Some(Box::new(HelloWorldFilter))
    }
}

struct HelloWorldFilter;

impl<'a> HelloWorldFilter {
}

impl Context for HelloWorldFilter {
    fn on_http_call_response(&mut self, _: u32, _: usize, _: usize, _: usize) {
        if let Some(status) = self.get_http_call_response_header(":status") {
            if status == "200" {
                info!("Access granted.");
                self.resume_http_request();
                return;
            }
        }

        info!("Access forbidden.");

        self.send_http_response(
            403,
            vec![("Powered-By", "proxy-wasm")],
            Some(b"Access forbidden.\n"),
        );
    }
}

impl HttpContext for HelloWorldFilter {
    fn on_http_request_headers(&mut self, _: usize) -> Action {
        info!("On http request.");

        match self.get_http_request_header("authorization") {
            Some(authorization) => {
                self.dispatch_http_call(
                    "httpbin",
                    vec![
                        (":method", "GET"),
                        (":path", "/bearer"),
                        (":authority", "httpbin.org"),
                        ("authorization", &authorization)
                    ],
                    None,
                    vec![],
                    Duration::from_secs(5),
                ).unwrap();
                Action::Pause
            }
            _ => {
                info!("Not found authorization header.");
                self.send_http_response(
                    403,
                    vec![],
                    Some(b"Access forbidden.\n"),
                );
                Action::Pause
            }
        }
    }

    fn on_http_response_headers(&mut self, _: usize) -> Action {
        self.set_http_response_header("Powered-By", Some("proxy-wasm"));
        Action::Continue
    }
}
