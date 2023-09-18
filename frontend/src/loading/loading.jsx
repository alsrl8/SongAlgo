import React from "react";
import {ConfigProvider, Modal, Spin} from "antd";

const Loading = ({isLoading, loadingText}) => {
    // const fetchData = async () => {
    //     setLoading(true);
    //     setTimeout(() => {
    //         setLoading(false);
    //     }, 10 * 1000);
    // };
    //
    // React.useEffect(() => {
    //     fetchData();
    // }, []);

    return (
        <ConfigProvider
            theme={{
                components: {
                    Spin: {
                        colorPrimary: "#aaaaaa",
                        dotSizeLG: 128,
                    },
                    Modal: {
                        contentBg: "transparent",
                        boxShadow: "0",
                    }
                },
            }}>
            <Modal
                open={isLoading}
                footer={null}
                closable={false}
                centered={true}
            >
                <div style={{textAlign: "center", display: "flex", flexDirection: "column"}}>
                    <Spin spinning={isLoading} size="large"/>
                    <h1>{loadingText}</h1>
                </div>
            </Modal>
        </ConfigProvider>
    );
}

export default Loading;
