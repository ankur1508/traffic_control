/*
 * Copyright 2015 Comcast Cable Communications Management, LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package com.comcast.cdn.traffic_control.traffic_router.core;

import org.apache.log4j.PropertyConfigurator;
import org.apache.wicket.util.time.Duration;
import org.eclipse.jetty.server.Server;
import org.eclipse.jetty.server.bio.SocketConnector;
import org.eclipse.jetty.webapp.WebAppContext;

public class TrafficRouterStart {
//	private static final Logger LOGGER = Logger.getLogger(TrafficRouterStart.class);

	public static void main(String[] args) throws Exception {
		PropertyConfigurator.configure("src/test/resources/log4j.properties");

		int timeout = (int) Duration.ONE_HOUR.getMilliseconds();

		Server server = new Server();
		SocketConnector connector = new SocketConnector();

		// Set some timeout options to make debugging easier.
		connector.setMaxIdleTime(timeout);
		connector.setSoLingerTime(-1);
		connector.setPort(8081);
		server.addConnector(connector);

		// check if a keystore for a SSL certificate is available, and
		// if so, start a SSL connector on port 8443. By default, the
		// quickstart comes with a Apache Wicket Quickstart Certificate
		// that expires about half way september 2021. Do not use this
		// certificate anywhere important as the passwords are available
		// in the source.

		//        Resource keystore = Resource.newClassPathResource("/keystore");
		//        if (keystore != null && keystore.exists()) {
		//            connector.setConfidentialPort(8443);
		//
		//            SslContextFactory factory = new SslContextFactory();
		//            factory.setKeyStoreResource(keystore);
		//            factory.setKeyStorePassword("wicket");
		//            factory.setTrustStoreResource(keystore);
		//            factory.setKeyManagerPassword("wicket");
		//            SslSocketConnector sslConnector = new SslSocketConnector(factory);
		//            sslConnector.setMaxIdleTime(timeout);
		//            sslConnector.setPort(8443);
		//            sslConnector.setAcceptors(4);
		//            server.addConnector(sslConnector);
		//
		//            System.out.println("SSL access to the quickstart has been enabled on port 8443");
		//            System.out.println("You can access the application using SSL on https://localhost:8443");
		//            System.out.println();
		//        }

		WebAppContext bb = new WebAppContext();
		bb.setServer(server);
		bb.setContextPath("/");
		bb.setWar("src/main/webapp");

		// START JMX SERVER
		// MBeanServer mBeanServer = ManagementFactory.getPlatformMBeanServer();
		// MBeanContainer mBeanContainer = new MBeanContainer(mBeanServer);
		// server.getContainer().addEventListener(mBeanContainer);
		// mBeanContainer.start();

		server.setHandler(bb);

		try {
			System.out.println(">>> STARTING EMBEDDED JETTY SERVER, PRESS ANY KEY TO STOP");
			server.start();
			System.in.read();
			System.out.println(">>> STOPPING EMBEDDED JETTY SERVER");
			server.stop();
			server.join();
		} catch (Exception e) {
			e.printStackTrace();
			System.exit(1);
		}
	}
}
