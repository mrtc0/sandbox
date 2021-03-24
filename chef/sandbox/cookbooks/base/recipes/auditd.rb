# frozen_string_literal: true

# Cookbook Name:: base
# Recipe:: auditd

include_recipe 'auditd'
auditd_ruleset node['auditd']['ruleset']
