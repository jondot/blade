require 'FileUtils'
include FileUtils

def error_if(condition, msg)
  return unless condition
  puts msg
  exit(1)
end


tests = %w{ interpolation }

tests.each do |test|
  test_dir = File.expand_path("fixtures/full_integration/#{test}", File.dirname(__FILE__))
  cd test_dir
  rm_rf "after"
  cp_r "before", "after"
  blade_out = `blade`
  error_if(!blade_out.empty?, blade_out)
  exit(1) unless blade_out.empty?
  diff_out = `git diff #{test_dir}`
  error_if(!diff_out.empty?, diff_out)
end





